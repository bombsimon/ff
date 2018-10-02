package ff

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Match represents a matched file or directory.
type Match struct {
	Path     string
	Filename string
}

// FilesFromPattern will take a base directory to scan, a pattern to match and
// a recursive flag.
func FilesFromPattern(base, matchPattern, ignorePattern string, recursive bool) ([]Match, error) {
	var matchedFiles []Match

	err := filepath.Walk(base, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Don't be recursive if not told by user except for given base.
		if info.IsDir() && !recursive && !pathIsSame(info.Name(), base) {
			return filepath.SkipDir
		}

		// Skip hidden files and directories except current path.
		if strings.HasPrefix(info.Name(), ".") && !pathIsSame(info.Name(), base) {
			if info.IsDir() {
				return filepath.SkipDir
			}

			return nil
		}

		// Check if the file or dir matches the glob matchPattern.
		matches, err := filepath.Match(matchPattern, info.Name())
		if err != nil {
			return err
		}

		if matches {
			// If it matches, ensure that it doesn't also match the ignorePattern.
			negativeMatches, err := filepath.Match(ignorePattern, info.Name())
			if err != nil {
				return err
			}

			if negativeMatches {
				return nil
			}

			dir, file := filepath.Split(path)
			if dir == "" {
				dir = "./"
			}

			matchedFiles = append(matchedFiles, Match{dir, file})
		}

		return nil
	})

	if err != nil {
		return matchedFiles, err
	}

	return matchedFiles, nil
}

// FullName returns the full path plus name for a file.
func (m *Match) FullName() string {
	return filepath.Join(m.Path, m.Filename)
}

// GetFilesFromDir returns all files from given directory.
func GetFilesFromDir(dir string) ([]Match, error) {
	return parseDir(dir, true, false)
}

// GetDirsFromDir returns all directories from given directory.
func GetDirsFromDir(dir string) ([]Match, error) {
	return parseDir(dir, false, true)
}

// GetAllFromDir returns all files and directories from a given directory.
func GetAllFromDir(dir string) ([]Match, error) {
	return parseDir(dir, true, true)
}

func parseDir(dir string, files, dirs bool) ([]Match, error) {
	var m []Match

	dirContent, err := ioutil.ReadDir(dir)
	if err != nil {
		return m, err
	}

	for _, info := range dirContent {
		if strings.HasPrefix(info.Name(), ".") {
			continue
		}

		switch info.IsDir() {
		case true:
			if dirs {
				fullDir := filepath.Join(dir, info.Name())
				m = append(m, Match{fullDir, ""})
			}
		case false:
			if files {
				m = append(m, Match{dir, info.Name()})
			}
		}
	}

	return m, nil
}

func pathIsSame(p1, p2 string) bool {
	return filepath.Clean(p1) == filepath.Clean(p2)
}

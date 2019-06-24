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
func FilesFromPattern(pattern string) ([]string, error) {
	// If double glob pattern isn't used use original filepath glob.
	if !strings.Contains(pattern, "**") {
		return filepath.Glob(pattern)
	}

	return expandGlobs(pattern)
}

func expandGlobs(pattern string) ([]string, error) {
	var (
		matches = []string{""}
		globs   = strings.Split(pattern, "**")
	)

	for _, glob := range globs {
		var (
			foundFiles = map[string]struct{}{}
		)

		for _, match := range matches {
			paths, err := filepath.Glob(match + glob)
			if err != nil {
				return nil, err
			}

			for _, path := range paths {
				err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
					if err != nil {
						return err
					}

					foundFiles[path] = struct{}{}

					return nil
				})

				if err != nil {
					return nil, err
				}
			}
		}

		matches = []string{}

		for match := range foundFiles {
			matches = append(matches, match)
		}
	}

	// Fix up return value for nil input.
	if globs == nil && len(matches) > 0 && matches[0] == "" {
		matches = matches[1:]
	}

	return matches, nil
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

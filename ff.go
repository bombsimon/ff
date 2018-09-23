package ff

import (
	"fmt"
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

	// Support go package syntax where './...' means '.' and allt it's sub
	// directories. When this match pattern is given we look for all Go files
	// recursively.
	if matchPattern == "./..." {
		matchPattern = "*.go"
		recursive = true
	}

	matchedFiles, err := GetAllFiles(base, matchPattern, ignorePattern, recursive, matchedFiles)
	if err != nil {
		return matchedFiles, err
	}

	return matchedFiles, nil
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

// GetAllFiles returns a slice of matches for all files found in given path.
// Without considering any matching.
func GetAllFiles(base string, match, ignore string, recursive bool, m []Match) ([]Match, error) {
	fmt.Printf("Checking for %s (ignoring %s) in %s, recursive is %v\n", match, ignore, base, recursive)

	matchGlob := filepath.Join("./", base, match)
	ignoreGlob := filepath.Join("./", base, ignore)

	matches, err := filepath.Glob(matchGlob)
	if err != nil {
		return m, err
	}

	ignores, err := filepath.Glob(ignoreGlob)
	if err != nil {
		return m, err
	}

	ignoreMap := mapFromSlice(ignores)

	// Add all files from match pattern (and not in ignore pattern).
	for _, file := range matches {
		// Skip file if ignore pattern.
		if _, ok := ignoreMap[file]; ok {
			continue
		}

		info, err := os.Stat(file)
		if err != nil {
			return m, err
		}

		// We don't see matched folders as file matches.
		if info.IsDir() {
			continue
		}

		m = append(m, Match{base, info.Name()})
	}

	if recursive {
		subdirs, err := GetDirsFromDir(base)
		if err != nil {
			return m, err
		}

		for _, subdir := range subdirs {
			subM, err := GetAllFiles(subdir.Path, match, ignore, recursive, m)
			if err != nil {
				return m, err
			}

			fmt.Println("should add", subM)
		}
	}

	return m, nil
}

func mapFromSlice(s []string) map[string]struct{} {
	uniq := make(map[string]struct{})

	for _, v := range s {
		uniq[v] = struct{}{}
	}

	return uniq
}

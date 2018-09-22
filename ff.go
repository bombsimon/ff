package ff

import (
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
func FilesFromPattern(base, matchPattern string, recursive bool) ([]Match, error) {
	var matchedFiles []Match

	// Support go package syntax where './...' means '.' and allt it's sub
	// directories. When this match pattern is given we look for all Go files
	// recursively.
	if matchPattern == "./..." {
		matchPattern = "*.go"
		recursive = true
	}

	files, err := filepath.Glob(matchPattern)
	if err != nil {
		return matchedFiles, err
	}

	/// If no expansion was made and only one file was found the input was that
	//exact file - just return it.
	if len(files) == 1 && files[0] == matchPattern {
		filePath, file := filepath.Split(matchPattern)

		matchedFiles = append(matchedFiles, Match{filePath, file})

		return matchedFiles, nil
	}

	for _, file := range files {
		info, err := os.Stat(file)
		if err != nil {
			return matchedFiles, err
		}

		if info.IsDir() {
		}
	}

	return []Match{}, nil
}

// GetAllFiles returns a slice of matches for all files found in given path.
// Without considering any matching.
func GetAllFiles(base string) ([]Match, error) {
	var m []Match

	err := filepath.Walk(base, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Ignore .git directory
		if strings.HasPrefix(path, ".git/") {
			return nil
		}

		// Ignore directories
		if info.IsDir() {
			return nil
		}

		filePath, file := filepath.Split(path)
		m = append(m, Match{filePath, file})

		return nil
	})

	return m, err
}

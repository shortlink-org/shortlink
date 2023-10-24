package tool

import (
	"os"
	"path/filepath"
)

func isExist(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}

	return false
}

func TrimQuotes(s string) string {
	if len(s) >= 2 { // nolint:gomnd
		if s[0] == '"' && s[len(s)-1] == '"' {
			return s[1 : len(s)-1]
		}
	}

	return s
}

func GetDirectories(root string, skipDirs []string) ([]string, error) {
	dirs := []string{}

	if err := filepath.Walk(
		root,
		func(path string, f os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			if f.IsDir() && isExist(skipDirs, f.Name()) {
				return filepath.SkipDir
			}

			if f.IsDir() {
				dirs = append(dirs, path)
			}

			return nil
		},
	); err != nil {
		return nil, err
	}

	return dirs, nil
}

func SaveToFile(filename, payload string) error {
	file, err := os.Create(filename) // #nosec
	if err != nil {
		return err
	}

	defer func() {
		_ = file.Close()
	}()

	_, err = file.WriteString(payload)
	if err != nil {
		return err
	}

	return file.Sync()
}

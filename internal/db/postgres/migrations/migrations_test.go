package migrations

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckMigration(t *testing.T) {
	files, err := getFilesSQL()
	assert.Nil(t, err)

	regexList := map[string]int{}

	for _, file := range files {
		s := strings.Split(file, "_")
		regexList[s[0]] += 1
	}

	for key, count := range regexList {
		if count == 1 || count > 2 {
			assert.Fail(t, fmt.Sprintf("migration with number %s as %d, but expect 0 or 2", key, count))
		}
	}
}

func getFilesSQL() ([]string, error) {
	files := []string{}
	err := filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if filepath.Ext(path) == ".sql" {
			files = append(files, path)
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

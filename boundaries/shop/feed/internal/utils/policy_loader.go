package utils

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"

	"github.com/shortlink-org/shortlink/boundaries/shop/feed/internal/usecase"
)

func LoadPolicy(filePath string) (usecase.Policy, error) {
	var policy usecase.Policy
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return policy, fmt.Errorf("error reading policy: %w", err)
	}
	err = yaml.Unmarshal(data, &policy)
	if err != nil {
		return policy, fmt.Errorf("error unmarshaling YAML: %w", err)
	}
	return policy, nil
}

func GetPolicyFiles(policyDir string) ([]string, error) {
	files, err := filepath.Glob(filepath.Join(policyDir, "*.yaml"))
	if err != nil {
		return nil, fmt.Errorf("error getting policy files: %w", err)
	}
	return files, nil
}

//go:build unit

package main_test

import (
	"os/exec"
	"strings"
	"testing"

	"github.com/shortlink-org/shortlink/pkg/fsroot"
)

func TestMongoORMGeneration(t *testing.T) {
	// Path to the proto file
	protoPath := "fixtures/link.proto"

	// Running protoc with the go-orm plugin and postgres flag
	cmd := exec.Command("protoc",
		"--go-orm_out=./output",
		"--go-orm_opt=orm=mongo,pkg=example,filter=Link;Car;Bus,orm=skip",
		"--proto_path=.",
		protoPath,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("protoc failed: %s, %v", string(output), err)
	}

	// Create SafeFS rooted at the output directory to safely read generated files
	fs, err := fsroot.NewSafeFS("./output")
	if err != nil {
		t.Fatalf("Failed to create SafeFS for output directory: %v", err)
	}
	defer fs.Close()

	// Check if the output file exists and contains MongoDB-specific ORM code
	data, err := fs.ReadFile("link.mongo.orm.go")
	if err != nil {
		t.Fatalf("Failed to read generated file: %v", err)
	}

	// Examples of PostgreSQL-specific checks you might perform
	expectedContents := []string{
		"\"go.mongodb.org/mongo-driver/v2/bson\"", // Check for PostgreSQL specific library import
		// Add more PostgreSQL-specific code snippets to check for
	}

	for _, content := range expectedContents {
		if !strings.Contains(string(data), content) {
			t.Errorf("Generated file does not contain expected PostgreSQL content: %s", content)
		}
	}
}

// func TestFilter_BuildMongoFilter(t *testing.T) {
// 	tests := []struct {
// 		name     string
// 		filter   *example.FilterLink
// 		expected bson.M
// 	}{
// 		{
// 			name: "Test Url Contains",
// 			filter: example.NewFilter(&domain.FilterLink{
// 				Url: &domain.StringFilterInput{Contains: []string{"example.com"}},
// 			}),
// 			expected: bson.M{"url": bson.M{"$in": []string{"example.com"}}},
// 		},
// 		{
// 			name: "Hash Equals and Describe NotContains",
// 			filter: example.NewFilter(&domain.FilterLink{
// 				Hash:     &domain.StringFilterInput{Eq: "123abc"},
// 				Describe: &domain.StringFilterInput{NotContains: []string{"test"}},
// 			}),
// 			expected: bson.M{
// 				"hash":     bson.M{"$eq": "123abc"},
// 				"describe": bson.M{"$nin": []string{"test"}},
// 			},
// 		},
// 		// Add more test cases for other conditions...
// 	}
//
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			actual := tt.filter.BuildMongoFilter()
// 			require.Equal(t, tt.expected, actual, "Mongo filter does not match expected")
// 		})
// 	}
// }

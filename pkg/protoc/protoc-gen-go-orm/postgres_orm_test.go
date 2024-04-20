//go:build unit

package main_test

import (
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestPostgresORMGeneration(t *testing.T) {
	// Path to the proto file
	protoPath := "fixtures/link.proto"

	// Running protoc with the go-orm plugin and postgres flag
	cmd := exec.Command("protoc",
		"--go-orm_out=./output",
		"--go-orm_opt=orm=postgres,pkg=example,filter=Link,common_path=github.com/shortlink-org/shortlink/boundaries/link/link/domain/link/v1",
		"--proto_path=.",
		protoPath,
	)

	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("protoc failed: %s, %v", string(output), err)
	}

	// Check if the output file exists and contains PostgreSQL-specific ORM code
	// You would specify the expected output filename based on your plugin's file naming scheme
	expectedFile := "./output/link.postgres.orm.go"
	data, err := os.ReadFile(expectedFile)
	if err != nil {
		t.Fatalf("Failed to read generated file: %v", err)
	}

	// Examples of PostgreSQL-specific checks you might perform
	expectedContents := []string{
		"\"github.com/Masterminds/squirrel\"", // Check for PostgreSQL specific library import
		// Add more PostgreSQL-specific code snippets to check for
	}

	for _, content := range expectedContents {
		if !strings.Contains(string(data), content) {
			t.Errorf("Generated file does not contain expected PostgreSQL content: %s", content)
		}
	}
}

// func TestFilterLink_BuildFilter(t *testing.T) {
// 	tests := []struct {
// 		name         string
// 		filter       *example.FilterLink
// 		expectedSQL  string
// 		expectedArgs []any
// 	}{
// 		{
// 			name: "Equal and Contains",
// 			filter: example.NewFilter(&domain.FilterLink{
// 				Url: &domain.StringFilterInput{
// 					Eq: "https://example.com",
// 				},
// 				Describe: &domain.StringFilterInput{
// 					Contains: []string{"test", "other"},
// 				},
// 			}),
// 			expectedSQL:  "SELECT * FROM links WHERE url = ? AND (describe LIKE ? OR describe LIKE ?)",
// 			expectedArgs: []any{"https://example.com", "%test%", "%other%"},
// 		},
// 		{
// 			name: "Not Equal and Starts With",
// 			filter: example.NewFilter(&domain.FilterLink{
// 				Url: &domain.StringFilterInput{
// 					Ne: "https://example.org",
// 				},
// 				Describe: &domain.StringFilterInput{
// 					StartsWith: "start",
// 				},
// 			}),
// 			expectedSQL:  "SELECT * FROM links WHERE url <> ? AND describe LIKE '%' || ?",
// 			expectedArgs: []any{"https://example.org", "start"},
// 		},
// 		{
// 			name: "Greater Than and Ends With",
// 			filter: example.NewFilter(&domain.FilterLink{
// 				Url: &domain.StringFilterInput{
// 					Gt: "https://example.com/a",
// 				},
// 				Describe: &domain.StringFilterInput{
// 					EndsWith: "end",
// 				},
// 			}),
// 			expectedSQL:  "SELECT * FROM links WHERE url > ? AND describe LIKE ? || '%'",
// 			expectedArgs: []any{"https://example.com/a", "end"},
// 		},
// 		{
// 			name: "Less Than and Is Empty",
// 			filter: example.NewFilter(&domain.FilterLink{
// 				Url: &domain.StringFilterInput{
// 					Lt: "https://example.com/z",
// 				},
// 				Describe: &domain.StringFilterInput{
// 					IsEmpty: true,
// 				},
// 			}),
// 			expectedSQL:  "SELECT * FROM links WHERE url < ? AND describe = '' OR describe IS NULL",
// 			expectedArgs: []any{"https://example.com/z"},
// 		},
// 		{
// 			name: "Complex - Multiple Conditions",
// 			filter: example.NewFilter(&domain.FilterLink{
// 				Url: &domain.StringFilterInput{
// 					Ne:         "https://example.org",
// 					StartsWith: "https",
// 					EndsWith:   ".com",
// 				},
// 				Describe: &domain.StringFilterInput{
// 					Contains:    []string{"test"},
// 					NotContains: []string{"example"},
// 					Gt:          "a",
// 					Lt:          "m",
// 				},
// 			}),
// 			expectedSQL: "SELECT * FROM links WHERE url <> ? AND url LIKE '%' || ? AND url LIKE ? || '%' AND " +
// 				"describe < ? AND describe > ? AND (describe LIKE ?) AND (describe NOT LIKE ?)",
// 			expectedArgs: []any{
// 				"https://example.org", "https", ".com",
// 				"m", "a", "%test%", "%example%", // Adjusted to match actual behavior
// 			},
// 		},
// 	}
//
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			query := squirrel.Select("*").From("links")
// 			query = tt.filter.BuildFilter(query)
// 			sql, args, err := query.ToSql()
//
// 			require.NoError(t, err, "Failed to build SQL query")
// 			require.Equal(t, tt.expectedSQL, sql, "SQL query does not match expected")
// 			require.Equal(t, tt.expectedArgs, args, "Query arguments do not match expected")
// 		})
// 	}
// }

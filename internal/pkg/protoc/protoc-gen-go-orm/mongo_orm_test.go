//go:build unit

package main

import (
	"testing"

	"github.com/stretchr/testify/require"
	"go.mongodb.org/mongo-driver/bson"

	"github.com/shortlink-org/shortlink/internal/pkg/protoc/protoc-gen-go-orm/fixtures"
)

func TestFilter_BuildMongoFilter(t *testing.T) {
	tests := []struct {
		name     string
		filter   fixtures.FilterLink
		expected bson.M
	}{
		{
			name: "Test Url Contains",
			filter: fixtures.FilterLink{
				Url: &fixtures.StringFilterInput{Contains: []string{"example.com"}},
			},
			expected: bson.M{"url": bson.M{"$in": []string{"example.com"}}},
		},
		{
			name: "Hash Equals and Describe NotContains",
			filter: fixtures.FilterLink{
				Hash:     &fixtures.StringFilterInput{Eq: "123abc"},
				Describe: &fixtures.StringFilterInput{NotContains: []string{"test"}},
			},
			expected: bson.M{
				"hash":     bson.M{"$eq": "123abc"},
				"describe": bson.M{"$nin": []string{"test"}},
			},
		},
		// Add more test cases for other conditions...
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := tt.filter.BuildMongoFilter()
			require.Equal(t, tt.expected, actual, "Mongo filter does not match expected")
		})
	}
}

package main

import (
	"bytes"
	"encoding/xml"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/shortlink-org/shortlink/boundaries/shop/feed/internal/infrastructure/persistence"
	"github.com/shortlink-org/shortlink/boundaries/shop/feed/internal/interfaces/controller"
	"github.com/shortlink-org/shortlink/boundaries/shop/feed/internal/usecase"
)

func TestGenerateFeeds(t *testing.T) {
	// Ensure the output directory exists and is empty
	outDir := "out"
	err := os.RemoveAll(outDir)
	require.NoError(t, err)
	err = os.MkdirAll(outDir, os.ModePerm)
	require.NoError(t, err)

	// Set up the dependencies
	goodsRepo := persistence.NewGoodsJSONRepository("tests/fixtures/phone.json")
	goodsUseCase := usecase.NewGoodsUseCase(goodsRepo)
	goodsController := controller.NewGoodsController(goodsUseCase)

	// Call the GenerateFeeds function
	err = goodsController.GenerateFeeds("policy", outDir)
	require.NoError(t, err)

	// Get the list of expected feed files from ./tests/dump/
	expectedFiles, err := filepath.Glob("tests/dump/*.xml")
	require.NoError(t, err)

	for _, expectedFile := range expectedFiles {
		// Get the base name of the file
		fileName := filepath.Base(expectedFile)

		// Read the expected content
		expectedContent, err := os.ReadFile(expectedFile)
		require.NoError(t, err)
		expectedContent, err = normalizeXML(expectedContent)
		require.NoError(t, err)

		// Read the generated content
		generatedFile := filepath.Join(outDir, fileName)
		generatedContent, err := os.ReadFile(generatedFile)
		require.NoError(t, err, "Generated file %s does not exist", generatedFile)
		generatedContent, err = normalizeXML(generatedContent)
		require.NoError(t, err)

		// Compare the contents
		require.Equal(t, string(expectedContent), string(generatedContent), "Contents of %s do not match", fileName)
	}
}

func normalizeXML(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	decoder := xml.NewDecoder(bytes.NewReader(data))
	encoder := xml.NewEncoder(&buf)
	for {
		token, err := decoder.Token()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil, err
		}
		err = encoder.EncodeToken(token)
		if err != nil {
			return nil, err
		}
	}
	encoder.Flush()
	return buf.Bytes(), nil
}

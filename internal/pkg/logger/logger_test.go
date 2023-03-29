//go:build unit
// +build unit

package logger

import (
	"bytes"
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/segmentio/encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/goleak"

	"github.com/shortlink-org/shortlink/internal/pkg/logger/field"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

// TestOutputInfoWithContextZap ...
func TestOutputInfoWithContextZap(t *testing.T) {
	var b bytes.Buffer

	conf := Configuration{
		Level:      INFO_LEVEL,
		Writer:     &b,
		TimeFormat: time.RFC822,
	}

	log, err := NewLogger(Zap, conf)
	require.NoError(t, err, "Error init a logger")

	log.InfoWithContext(context.Background(), "Hello World")

	expectedTime := time.Now().Format(time.RFC822)
	expected := map[string]interface{}{
		"level":     "info",
		"timestamp": expectedTime,
		"caller":    "logger/logger_test.go:38",
		"msg":       "Hello World",
		"traceID":   "00000000000000000000000000000000",
	}
	var response map[string]interface{}
	require.NoError(t, json.Unmarshal(b.Bytes(), &response), "Error unmarshalling")

	if !reflect.DeepEqual(expected, response) {
		assert.Equal(t, expected, response)
	}
}

func BenchmarkOutputZap(bench *testing.B) {
	var b bytes.Buffer

	conf := Configuration{
		Level:      INFO_LEVEL,
		Writer:     &b,
		TimeFormat: time.RFC822,
	}

	log, _ := NewLogger(Zap, conf)

	for i := 0; i < bench.N; i++ {
		log.Info("Hello World")
	}
}

func TestOutputInfoWithContextLogrus(t *testing.T) {
	var b bytes.Buffer

	conf := Configuration{
		Level:      INFO_LEVEL,
		Writer:     &b,
		TimeFormat: time.RFC822,
	}

	log, err := NewLogger(Logrus, conf)
	require.NoError(t, err, "Error init a logger")

	log.InfoWithContext(context.Background(), "Hello World")

	expectedTime := time.Now().Format(time.RFC822)
	expected := map[string]interface{}{
		"level":     "info",
		"timestamp": expectedTime,
		"msg":       "Hello World",
		"traceID":   "00000000000000000000000000000000",
	}
	var response map[string]interface{}
	require.NoError(t, json.Unmarshal(b.Bytes(), &response), "Error unmarshalling")
	assert.Equal(t, expected, response)
}

func BenchmarkOutputLogrus(bench *testing.B) {
	var b bytes.Buffer

	conf := Configuration{
		Level:      INFO_LEVEL,
		Writer:     &b,
		TimeFormat: time.RFC822,
	}

	log, _ := NewLogger(Logrus, conf)

	for i := 0; i < bench.N; i++ {
		log.Info("Hello World")
	}
}

func TestFieldsZap(t *testing.T) {
	var b bytes.Buffer

	conf := Configuration{
		Level:      INFO_LEVEL,
		Writer:     &b,
		TimeFormat: time.RFC822,
	}

	log, err := NewLogger(Zap, conf)
	require.NoError(t, err, "Error init a logger")

	log.InfoWithContext(context.Background(), "Hello World", field.Fields{
		"hello": "world",
		"first": 1,
	})

	expectedTime := time.Now().Format(time.RFC822)
	expected := map[string]interface{}{
		"level":     "info",
		"timestamp": expectedTime,
		"msg":       "Hello World",
		"caller":    "logger/logger_test.go:126",
		"first":     float64(1),
		"hello":     "world",
		"traceID":   "00000000000000000000000000000000",
	}
	var response map[string]interface{}
	require.NoError(t, json.Unmarshal(b.Bytes(), &response), "Error unmarshalling")

	if !reflect.DeepEqual(expected, response) {
		assert.Equal(t, expected, response)
	}
}

func TestFieldsLogrus(t *testing.T) {
	var b bytes.Buffer

	conf := Configuration{
		Level:      INFO_LEVEL,
		Writer:     &b,
		TimeFormat: time.RFC822,
	}

	log, err := NewLogger(Logrus, conf)
	require.NoError(t, err, "Error init a logger")

	log.Info("Hello World", field.Fields{
		"hello": "world",
		"first": 1,
	})

	expectedTime := time.Now().Format(time.RFC822)
	expected := map[string]interface{}{
		"level":     "info",
		"timestamp": expectedTime,
		"msg":       "Hello World",
		"first":     float64(1),
		"hello":     "world",
	}
	var response map[string]interface{}
	require.NoError(t, json.Unmarshal(b.Bytes(), &response), "Error unmarshalling")
	assert.Equal(t, expected, response)
}

func TestSetLevel(t *testing.T) {
	loggerList := []int{Zap, Logrus}

	for _, logger := range loggerList {
		var b bytes.Buffer

		conf := Configuration{
			Level:      FATAL_LEVEL,
			Writer:     &b,
			TimeFormat: time.RFC822,
		}

		log, err := NewLogger(logger, conf)
		require.NoError(t, err, "Error init a logger")

		log.Info("Hello World")

		expectedStr := ``

		if b.String() != expectedStr {
			assert.Errorf(t, err, "Expected: %sgot: %s", expectedStr, b.String())
		}
	}
}

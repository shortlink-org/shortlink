//go:build unit
// +build unit

package logger

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"

	"github.com/batazor/shortlink/internal/pkg/logger/field"
)

func TestMain(m *testing.M) {
	goleak.VerifyTestMain(m)
}

// TestOutput ...
func TestOutputZap(t *testing.T) {
	var b bytes.Buffer

	conf := Configuration{
		Level:      INFO_LEVEL,
		Writer:     &b,
		TimeFormat: time.RFC822,
	}

	log, err := NewLogger(Zap, conf)
	assert.Nil(t, err, "Error init a logger")

	log.Info("Hello World")

	expectedTime := time.Now().Format(time.RFC822)
	expected := map[string]interface{}{
		"level":     "info",
		"timestamp": expectedTime,
		"caller":    "logger/logger_test.go:36",
		"msg":       "Hello World",
	}
	var response map[string]interface{}
	assert.Nil(t, json.Unmarshal(b.Bytes(), &response), "Error unmarshalling")

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

func TestOutputLogrus(t *testing.T) {
	var b bytes.Buffer

	conf := Configuration{
		Level:      INFO_LEVEL,
		Writer:     &b,
		TimeFormat: time.RFC822,
	}

	log, err := NewLogger(Logrus, conf)
	assert.Nil(t, err, "Error init a logger")

	log.Info("Hello World")

	expectedTime := time.Now().Format(time.RFC822)
	expected := map[string]interface{}{
		"level":     "info",
		"timestamp": expectedTime,
		"msg":       "Hello World",
	}
	var response map[string]interface{}
	assert.Nil(t, json.Unmarshal(b.Bytes(), &response), "Error unmarshalling")
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
	assert.Nil(t, err, "Error init a logger")

	log.Info("Hello World", field.Fields{
		"hello": "world",
		"first": 1,
	})

	expectedTime := time.Now().Format(time.RFC822)
	expected := map[string]interface{}{
		"level":     "info",
		"timestamp": expectedTime,
		"msg":       "Hello World",
		"caller":    "logger/logger_test.go:122",
		"first":     float64(1),
		"hello":     "world",
	}
	var response map[string]interface{}
	assert.Nil(t, json.Unmarshal(b.Bytes(), &response), "Error unmarshalling")

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
	assert.Nil(t, err, "Error init a logger")

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
	assert.Nil(t, json.Unmarshal(b.Bytes(), &response), "Error unmarshalling")
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
		assert.Nil(t, err, "Error init a logger")

		log.Info("Hello World")

		expectedStr := ``

		if b.String() != expectedStr {
			assert.Errorf(t, err, "Expected: %sgot: %s", expectedStr, b.String())
		}
	}
}

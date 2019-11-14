package logger

import (
	"bytes"
	"encoding/json"
	"reflect"
	"testing"
	"time"
)

// TestOutput ...
func TestOutputZap(t *testing.T) { //nolint unused
	var b bytes.Buffer

	conf := Configuration{
		Level:      INFO_LEVEL,
		Writer:     &b,
		TimeFormat: time.RFC822,
	}

	log, err := NewLogger(Zap, conf)
	if err != nil {
		t.Errorf("Error init a logger: %s", err)
	}

	log.Info("Hello World")

	expectedTime := time.Now().Format(time.RFC822)
	expected := map[string]interface{}{
		"@level":     "info",
		"@timestamp": expectedTime,
		"@caller":    "logger/logger_test.go:26",
		"@msg":       "Hello World",
	}
	var response map[string]interface{}
	if err := json.Unmarshal(b.Bytes(), &response); err != nil {
		t.Errorf("Error unmarshalling %s", err.Error())
	}

	if !reflect.DeepEqual(expected, response) {
		t.Errorf("Expected: %s\ngot: %s", expected, response)
	}
}

func TestOutputLogrus(t *testing.T) { //nolint unused
	var b bytes.Buffer

	conf := Configuration{
		Level:      INFO_LEVEL,
		Writer:     &b,
		TimeFormat: time.RFC822,
	}

	log, err := NewLogger(Logrus, conf)
	if err != nil {
		t.Errorf("Error init a logger: %s", err)
	}

	log.Info("Hello World")

	expectedTime := time.Now().Format(time.RFC822)
	expected := map[string]interface{}{
		"@level":     "info",
		"@timestamp": expectedTime,
		"@msg":       "Hello World",
	}
	var response map[string]interface{}
	if err := json.Unmarshal(b.Bytes(), &response); err != nil {
		t.Errorf("Error unmarshalling %s", err.Error())
	}

	if !reflect.DeepEqual(expected, response) {
		t.Errorf("Expected: %s\ngot: %s", expected, response)
	}
}

func TestFieldsZap(t *testing.T) { //nolint unused
	var b bytes.Buffer

	conf := Configuration{
		Level:      INFO_LEVEL,
		Writer:     &b,
		TimeFormat: time.RFC822,
	}

	log, err := NewLogger(Zap, conf)
	if err != nil {
		t.Errorf("Error init a logger: %s", err)
	}

	log.Info("Hello World", Fields{
		"hello": "world",
		"first": 1,
	})

	expectedTime := time.Now().Format(time.RFC822)
	expected := map[string]interface{}{
		"@level":     "info",
		"@timestamp": expectedTime,
		"@msg":       "Hello World",
		"@caller":    "logger/logger_test.go:91",
		"first":      float64(1),
		"hello":      "world",
	}
	var response map[string]interface{}
	if err := json.Unmarshal(b.Bytes(), &response); err != nil {
		t.Errorf("Error unmarshalling %s", err.Error())
	}

	if !reflect.DeepEqual(expected, response) {
		t.Errorf("Expected: %s\ngot: %s", expected, response)
	}
}

func TestFieldsLogrus(t *testing.T) { //nolint unused
	var b bytes.Buffer

	conf := Configuration{
		Level:      INFO_LEVEL,
		Writer:     &b,
		TimeFormat: time.RFC822,
	}

	log, err := NewLogger(Logrus, conf)
	if err != nil {
		t.Errorf("Error init a logger: %s", err)
	}

	log.Info("Hello World", Fields{
		"hello": "world",
		"first": 1,
	})

	expectedTime := time.Now().Format(time.RFC822)
	expected := map[string]interface{}{
		"@level":     "info",
		"@timestamp": expectedTime,
		"@msg":       "Hello World",
		"first":      float64(1),
		"hello":      "world",
	}
	var response map[string]interface{}
	if err := json.Unmarshal(b.Bytes(), &response); err != nil {
		t.Errorf("Error unmarshalling %s", err.Error())
	}

	if !reflect.DeepEqual(expected, response) {
		t.Errorf("Expected: %s\ngot: %s", expected, response)
	}
}

func TestSetLevel(t *testing.T) { //nolint unused
	loggerList := []int{Zap, Logrus}

	for _, logger := range loggerList {
		var b bytes.Buffer

		conf := Configuration{
			Level:      FATAL_LEVEL,
			Writer:     &b,
			TimeFormat: time.RFC822,
		}

		log, err := NewLogger(logger, conf)
		if err != nil {
			t.Errorf("Error init a logger: %s", err)
		}

		log.Info("Hello World")

		expectedStr := ``

		if b.String() != expectedStr {
			t.Errorf("Expected: %sgot: %s", expectedStr, b.String())
		}
	}
}

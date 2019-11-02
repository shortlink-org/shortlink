package logger

import (
	"bytes"
	"strings"
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
	expectedStr := `{"@level":"info","@timestamp":"` + expectedTime + `","@caller":"logger/logger_test.go:25","@msg":"Hello World"}`

	if strings.TrimRight(b.String(), "\n") != expectedStr {
		t.Errorf("Expected: %sgot: %s", expectedStr, b.String())
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
	expectedStr := `{"@level":"info","@msg":"Hello World","@timestamp":"` + expectedTime + `"}`

	if strings.TrimRight(b.String(), "\n") != expectedStr {
		t.Errorf("Expected: %s\ngot: %s", expectedStr, b.String())
	}
}

func TestSetLevelZap(t *testing.T) { //nolint unused
	var b bytes.Buffer

	conf := Configuration{
		Level:      FATAL_LEVEL,
		Writer:     &b,
		TimeFormat: time.RFC822,
	}

	log, err := NewLogger(Zap, conf)
	if err != nil {
		t.Errorf("Error init a logger: %s", err)
	}

	log.Info("Hello World")

	expectedStr := ``

	if strings.TrimRight(b.String(), "\n") != expectedStr {
		t.Errorf("Expected: %sgot: %s", expectedStr, b.String())
	}
}

func TestSetLevelLogrus(t *testing.T) { //nolint unused
	var b bytes.Buffer

	conf := Configuration{
		Level:      FATAL_LEVEL,
		Writer:     &b,
		TimeFormat: time.RFC822,
	}

	log, err := NewLogger(Logrus, conf)
	if err != nil {
		t.Errorf("Error init a logger: %s", err)
	}

	log.Info("Hello World")

	expectedStr := ``

	if strings.TrimRight(b.String(), "\n") != expectedStr {
		t.Errorf("Expected: %sgot: %s", expectedStr, b.String())
	}
}

func TestFieldZap(t *testing.T) { //nolint unused
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
	expectedStr := `{"@level":"info","@timestamp":"` + expectedTime + `","@caller":"logger/logger_test.go:119","@msg":"Hello World","hello":"world","first":1}`

	if strings.TrimRight(b.String(), "\n") != expectedStr {
		t.Errorf("Expected: %s\ngot: %s", expectedStr, b.String())
	}
}

func TestFieldLogrus(t *testing.T) { //nolint unused
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
	expectedStr := `{"@level":"info","@msg":"Hello World","@timestamp":"` + expectedTime + `","first":1,"hello":"world"}`

	if strings.TrimRight(b.String(), "\n") != expectedStr {
		t.Errorf("Expected: %s\ngot: %s", expectedStr, b.String())
	}
}

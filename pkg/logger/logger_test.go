package logger

import (
	"bytes"
	"fmt"
	"testing"
)

// TestOutput ...
func TestOutput(t *testing.T) { //nolint unused
	var b bytes.Buffer

	conf := Configuration{
		Level:  INFO_LEVEL,
		Writer: &b,
	}

	log, err := NewLogger(Zap, conf)
	if err != nil {
		t.Errorf("Error init a logger: %s", err)
	}

	log.Info("Run HTTP-CHI API")

	fmt.Println(b.String())
	fmt.Println(b.String())
}

// Test input/output
// -> log.Info("Run HTTP-CHI API")
// -> {"level":"info","msg":"Run HTTP-CHI API","time":"2019-11-02T14:29:24+03:00"}

// Test setlevel
// setLevel error
// -> log.Info("Run HTTP-CHI API")
// -> ...
// setLevel info
// -> log.Info("Run HTTP-CHI API")
// -> {"level":"info","msg":"Run HTTP-CHI API","time":"2019-11-02T14:29:24+03:00"}

// Test fields
//var fields = logger.Fields{
//	"hello": "world",
//}
// -> log.Info("Run HTTP-CHI API", fields)
// -> {"level":"info","msg":"Run HTTP-CHI API","time":"2019-11-02T14:29:24+03:00", "hello": "world"}

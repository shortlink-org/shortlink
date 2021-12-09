package v1

import (
	"fmt"
	"os"
	"testing"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/spf13/pflag"
)

var opts = godog.Options{
	Format: "progress",
	Output: colors.Colored(os.Stdout),
}

var sqlResponse string

func init() {
	godog.BindCommandLineFlags("godog.", &opts) // godog v0.11.0 and later
}

func TestMain(m *testing.M) {
	pflag.Parse()
	opts.Paths = pflag.Args()

	status := godog.TestSuite{
		Name:                "godogs",
		ScenarioInitializer: InitializeScenario,
		Options:             &opts,
	}.Run()

	// Optional: Run `testing` package's logic besides godog.
	if st := m.Run(); st > status {
		status = st
	}

	os.Exit(status)
}

func expression(sql string) error {
	_, err := New(sql)

	// nil -> "nil"
	if err == nil {
		sqlResponse = ""
	} else {
		sqlResponse = err.Error()
	}

	return nil
}

func theResponse(response string) error {
	if response != sqlResponse {
		return fmt.Errorf("incorrect parse result. expect: %s, but get: %s", response, sqlResponse)
	}

	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^"([^"]*)" expression$`, expression)
	ctx.Step(`^the response "([^"]*)"$`, theResponse)
}

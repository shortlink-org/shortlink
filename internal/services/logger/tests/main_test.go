package tests

import (
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

func init() {
	godog.BindCommandLineFlags("godog.", &opts) // godog v0.11.0 and later
}

func TestMain(m *testing.M) {
	pflag.Parse()
	opts.Paths = pflag.Args()

	status := godog.TestSuite{
		Name:    "godogs",
		Options: &opts,
	}.Run()

	// Optional: Run `testing` package's logic besides godog.
	if st := m.Run(); st > status {
		status = st
	}

	os.Exit(status)
}

func getCorrectPayloadWithNotNilFieldURLHttpsgooglecom(arg1 string) error {
	return nil
}

func getNewAEventFromMQ() error {
	return nil
}

func getRandomPayload(arg1 string) error {
	return nil
}

func printErrorMessageIncorrectFormatPayloadErrorMessage() error {
	return nil
}

func printLinkURLHttpsgooglecom() error {
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^Get correct payload with not nil field URL {"([^"]*)":"https:\/\/google\.com"}\$$`, getCorrectPayloadWithNotNilFieldURLHttpsgooglecom)
	ctx.Step(`^get new a event from MQ$`, getNewAEventFromMQ)
	ctx.Step(`^Get random payload {"([^"]*)":""}\$$`, getRandomPayload)
	ctx.Step(`^Get random payload "([^"]*)"\$$`, getRandomPayload)
	ctx.Step(`^print error message: Incorrect format payload error message\$$`, printErrorMessageIncorrectFormatPayloadErrorMessage)
	ctx.Step(`^print link URL https:\/\/google\.com\$$`, printLinkURLHttpsgooglecom)
}

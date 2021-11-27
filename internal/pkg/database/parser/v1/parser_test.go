package v1

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

func getEmptyQuery() error {
	return godog.ErrPending
}

func getEmptyTableName() error {
	return godog.ErrPending
}

func weGetAnErrorMessage() error {
	return godog.ErrPending
}

func weWantToParseAnSQLExpression() error {
	return godog.ErrPending
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^Get empty query$`, getEmptyQuery)
	ctx.Step(`^Get empty table name$`, getEmptyTableName)
	ctx.Step(`^we get an error message$`, weGetAnErrorMessage)
	ctx.Step(`^we want to parse an SQL expression$`, weWantToParseAnSQLExpression)
}

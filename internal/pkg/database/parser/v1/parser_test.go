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

func dELETEFROMAExpression(arg1 int) error {
	return godog.ErrPending
}

func dELETEFROMAWHEREBExpression1(arg1 int) error {
	return godog.ErrPending
}

func dELETEFROMAWHEREBExpression2(arg1, arg2 int) error {
	return godog.ErrPending
}

func dELETEFROMAWHEREExpression(arg1 int) error {
	return godog.ErrPending
}

func dELETEFROMExpression() error {
	return godog.ErrPending
}

func expression(arg1 string) error {
	return godog.ErrPending
}

func iNSERTINTOABExpression1(arg1 int) error {
	return godog.ErrPending
}

func iNSERTINTOABExpression2(arg1 int) error {
	return godog.ErrPending
}

func iNSERTINTOABVALUESExpression1(arg1 int) error {
	return godog.ErrPending
}

func iNSERTINTOABVALUESExpression2(arg1 int) error {
	return godog.ErrPending
}

func iNSERTINTOABVALUESExpression3(arg1, arg2 int) error {
	return godog.ErrPending
}

func iNSERTINTOABcdVALUESExpression1(arg1, arg2, arg3, arg4, arg5, arg6, arg7 int) error {
	return godog.ErrPending
}

func iNSERTINTOABcdVALUESExpression2(arg1, arg2, arg3, arg4 int) error {
	return godog.ErrPending
}

func iNSERTINTOAExpression1(arg1 int) error {
	return godog.ErrPending
}

func iNSERTINTOAExpression2(arg1 int) error {
	return godog.ErrPending
}

func iNSERTINTOAVALUESExpression(arg1, arg2 int) error {
	return godog.ErrPending
}

func iNSERTINTOExpression() error {
	return godog.ErrPending
}

func sELECTAAsZBAsYCFROMBExpression(arg1 int) error {
	return godog.ErrPending
}

func sELECTACDFROMBExpression(arg1 int) error {
	return godog.ErrPending
}

func sELECTACDFROMBWHEREAANDBExpression(arg1, arg2 int) error {
	return godog.ErrPending
}

func sELECTACDFROMBWHEREABExpression(arg1 int) error {
	return godog.ErrPending
}

func sELECTACDFROMBWHEREAExpression1(arg1, arg2 int) error {
	return godog.ErrPending
}

func sELECTACDFROMBWHEREAExpression2(arg1, arg2 int) error {
	return godog.ErrPending
}

func sELECTACDFROMBWHEREAExpression3(arg1 int) error {
	return godog.ErrPending
}

func sELECTACDFROMBWHEREAExpression4(arg1, arg2 int) error {
	return godog.ErrPending
}

func sELECTACDFROMBWHEREAExpression5(arg1 int) error {
	return godog.ErrPending
}

func sELECTACDFROMBWHEREAExpression6(arg1, arg2 int) error {
	return godog.ErrPending
}

func sELECTACDFROMBWHEREAExpression7(arg1, arg2 int) error {
	return godog.ErrPending
}

func sELECTACDFROMBWHEREExpression(arg1 int) error {
	return godog.ErrPending
}

func sELECTAFROMBExpression1() error {
	return godog.ErrPending
}

func sELECTAFROMBExpression2() error {
	return godog.ErrPending
}

func sELECTBFROMAExpression() error {
	return godog.ErrPending
}

func sELECTExpression() error {
	return godog.ErrPending
}

func sELECTFROMAExpression() error {
	return godog.ErrPending
}

func sELECTFROMBExpression() error {
	return godog.ErrPending
}

func selectAFRoMCExpression() error {
	return godog.ErrPending
}

func theResponseAtINSERTINTOExpectedAtLeastOneFieldToInsert() error {
	return godog.ErrPending
}

func theResponseAtINSERTINTONeedAtLeastOneRowToInsert() error {
	return godog.ErrPending
}

func theResponseAtINSERTINTOValueCountDoesntMatchFieldCount() error {
	return godog.ErrPending
}

func theResponseAtSELECTExpectedFieldToSELECT() error {
	return godog.ErrPending
}

func theResponseAtUPDATEExpected() error {
	return godog.ErrPending
}

func theResponseAtUPDATEExpectedQuotedValue() error {
	return godog.ErrPending
}

func theResponseAtWHEREConditionWithoutOperator() error {
	return godog.ErrPending
}

func theResponseAtWHEREEmptyWHEREClause() error {
	return godog.ErrPending
}

func theResponseAtWHEREWHEREClauseIsMandatoryForUPDATEDELETE() error {
	return godog.ErrPending
}

func theResponseNil() error {
	return godog.ErrPending
}

func theResponseQueryTypeCannotBeEmpty() error {
	return godog.ErrPending
}

func theResponseTableNameCannotBeEmpty() error {
	return godog.ErrPending
}

func uPDATEAExpression(arg1 int) error {
	return godog.ErrPending
}

func uPDATEASETBHelloCByeWHEREAANDBExpression(arg1, arg2, arg3 int) error {
	return godog.ErrPending
}

func uPDATEASETBHelloCByeWHEREAExpression(arg1, arg2 int) error {
	return godog.ErrPending
}

func uPDATEASETBHelloWHEREAExpression1(arg1 int) error {
	return godog.ErrPending
}

func uPDATEASETBHelloWHEREAExpression2(arg1, arg2 int) error {
	return godog.ErrPending
}

func uPDATEASETBHelloWHEREExpression(arg1 int) error {
	return godog.ErrPending
}

func uPDATEASETBHelloworldWHEREAExpression(arg1, arg2 int) error {
	return godog.ErrPending
}

func uPDATEASETBWHEREExpression1(arg1 int) error {
	return godog.ErrPending
}

func uPDATEASETBWHEREExpression2(arg1 int) error {
	return godog.ErrPending
}

func uPDATEASETExpression(arg1 int) error {
	return godog.ErrPending
}

func uPDATEExpression() error {
	return godog.ErrPending
}

func weGetTheQuery() error {
	return godog.ErrPending
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^DELETE FROM \'a(\d+)\' expression$`, dELETEFROMAExpression)
	ctx.Step(`^DELETE FROM \'a(\d+)\' WHERE b expression$`, dELETEFROMAWHEREBExpression1)
	ctx.Step(`^DELETE FROM \'a(\d+)\' WHERE b = \'(\d+)\' expression$`, dELETEFROMAWHEREBExpression2)
	ctx.Step(`^DELETE FROM \'a(\d+)\' WHERE expression$`, dELETEFROMAWHEREExpression)
	ctx.Step(`^DELETE FROM expression$`, dELETEFROMExpression)
	ctx.Step(`^"([^"]*)" expression$`, expression)
	ctx.Step(`^INSERT INTO \'a(\d+)\' \(b\) expression$`, iNSERTINTOABExpression1)
	ctx.Step(`^INSERT INTO \'a(\d+)\' \(b expression$`, iNSERTINTOABExpression2)
	ctx.Step(`^INSERT INTO \'a(\d+)\' \(b\) VALUES expression$`, iNSERTINTOABVALUESExpression1)
	ctx.Step(`^INSERT INTO \'a(\d+)\' \(b\) VALUES \( expression$`, iNSERTINTOABVALUESExpression2)
	ctx.Step(`^INSERT INTO \'a(\d+)\' \(b\) VALUES \(\'(\d+)\'\) expression$`, iNSERTINTOABVALUESExpression3)
	ctx.Step(`^INSERT INTO \'a(\d+)\' \(b,c,d\) VALUES \(\'(\d+)\',\'(\d+)\' ,\'(\d+)\'\),\(\'(\d+)\',\'(\d+)\',\'(\d+)\'\) expression$`, iNSERTINTOABcdVALUESExpression1)
	ctx.Step(`^INSERT INTO \'a(\d+)\' \(b,c,d\) VALUES \(\'(\d+)\',\'(\d+)\',\'(\d+)\'\) expression$`, iNSERTINTOABcdVALUESExpression2)
	ctx.Step(`^INSERT INTO \'a(\d+)\' \( expression$`, iNSERTINTOAExpression1)
	ctx.Step(`^INSERT INTO \'a(\d+)\' expression$`, iNSERTINTOAExpression2)
	ctx.Step(`^INSERT INTO \'a(\d+)\' \(\*\) VALUES \(\'(\d+)\'\) expression$`, iNSERTINTOAVALUESExpression)
	ctx.Step(`^INSERT INTO expression$`, iNSERTINTOExpression)
	ctx.Step(`^SELECT a as z, b as y, c FROM \'b(\d+)\' expression$`, sELECTAAsZBAsYCFROMBExpression)
	ctx.Step(`^SELECT a, c, d FROM \'b(\d+)\' expression$`, sELECTACDFROMBExpression)
	ctx.Step(`^SELECT a, c, d FROM \'b\' WHERE a != \'(\d+)\' AND b = \'(\d+)\' expression$`, sELECTACDFROMBWHEREAANDBExpression)
	ctx.Step(`^SELECT a, c, d FROM \'b(\d+)\' WHERE a != b expression$`, sELECTACDFROMBWHEREABExpression)
	ctx.Step(`^SELECT a, c, d FROM \'b(\d+)\' WHERE a >= \'(\d+)\' expression$`, sELECTACDFROMBWHEREAExpression1)
	ctx.Step(`^SELECT a, c, d FROM \'b(\d+)\' WHERE a > \'(\d+)\' expression$`, sELECTACDFROMBWHEREAExpression2)
	ctx.Step(`^SELECT a, c, d FROM \'b(\d+)\' WHERE a expression$`, sELECTACDFROMBWHEREAExpression3)
	ctx.Step(`^SELECT a, c, d FROM \'b(\d+)\' WHERE a <= \'(\d+)\' expression$`, sELECTACDFROMBWHEREAExpression4)
	ctx.Step(`^SELECT a, c, d FROM \'b(\d+)\' WHERE a = \'\' expression$`, sELECTACDFROMBWHEREAExpression5)
	ctx.Step(`^SELECT a, c, d FROM \'b(\d+)\' WHERE a != \'(\d+)\' expression$`, sELECTACDFROMBWHEREAExpression6)
	ctx.Step(`^SELECT a, c, d FROM \'b(\d+)\' WHERE a < \'(\d+)\' expression$`, sELECTACDFROMBWHEREAExpression7)
	ctx.Step(`^SELECT a, c, d FROM \'b(\d+)\' WHERE expression$`, sELECTACDFROMBWHEREExpression)
	ctx.Step(`^SELECT a, \* FROM \'b\' expression$`, sELECTAFROMBExpression1)
	ctx.Step(`^SELECT a FROM \'b\' expression$`, sELECTAFROMBExpression2)
	ctx.Step(`^SELECT b, FROM \'a\' expression$`, sELECTBFROMAExpression)
	ctx.Step(`^SELECT expression$`, sELECTExpression)
	ctx.Step(`^SELECT FROM \'a\' expression$`, sELECTFROMAExpression)
	ctx.Step(`^SELECT \* FROM \'b\' expression$`, sELECTFROMBExpression)
	ctx.Step(`^select a fRoM \'c\' expression$`, selectAFRoMCExpression)
	ctx.Step(`^the response at INSERT INTO: expected at least one field to insert$`, theResponseAtINSERTINTOExpectedAtLeastOneFieldToInsert)
	ctx.Step(`^the response at INSERT INTO: need at least one row to insert$`, theResponseAtINSERTINTONeedAtLeastOneRowToInsert)
	ctx.Step(`^the response at INSERT INTO: value count doesn\'t match field count$`, theResponseAtINSERTINTOValueCountDoesntMatchFieldCount)
	ctx.Step(`^the response at SELECT: expected field to SELECT$`, theResponseAtSELECTExpectedFieldToSELECT)
	ctx.Step(`^the response at UPDATE: expected \'=\'$`, theResponseAtUPDATEExpected)
	ctx.Step(`^the response at UPDATE: expected quoted value$`, theResponseAtUPDATEExpectedQuotedValue)
	ctx.Step(`^the response at WHERE: condition without operator$`, theResponseAtWHEREConditionWithoutOperator)
	ctx.Step(`^the response at WHERE: empty WHERE clause$`, theResponseAtWHEREEmptyWHEREClause)
	ctx.Step(`^the response at WHERE: WHERE clause is mandatory for UPDATE & DELETE$`, theResponseAtWHEREWHEREClauseIsMandatoryForUPDATEDELETE)
	ctx.Step(`^the response nil$`, theResponseNil)
	ctx.Step(`^the response query type cannot be empty$`, theResponseQueryTypeCannotBeEmpty)
	ctx.Step(`^the response table name cannot be empty$`, theResponseTableNameCannotBeEmpty)
	ctx.Step(`^UPDATE \'a(\d+)\' expression$`, uPDATEAExpression)
	ctx.Step(`^UPDATE \'a(\d+)\' SET b = \'hello\', c = \'bye\' WHERE a = \'(\d+)\' AND b = \'(\d+)\' expression$`, uPDATEASETBHelloCByeWHEREAANDBExpression)
	ctx.Step(`^UPDATE \'a(\d+)\' SET b = \'hello\', c = \'bye\' WHERE a = \'(\d+)\' expression$`, uPDATEASETBHelloCByeWHEREAExpression)
	ctx.Step(`^UPDATE \'a(\d+)\' SET b = \'hello\' WHERE a expression$`, uPDATEASETBHelloWHEREAExpression1)
	ctx.Step(`^UPDATE \'a(\d+)\' SET b = \'hello\' WHERE a = \'(\d+)\' expression$`, uPDATEASETBHelloWHEREAExpression2)
	ctx.Step(`^UPDATE \'a(\d+)\' SET b = \'hello\' WHERE expression$`, uPDATEASETBHelloWHEREExpression)
	ctx.Step(`^UPDATE \'a(\d+)\' SET b = \'hello\\\'world\' WHERE a = \'(\d+)\' expression$`, uPDATEASETBHelloworldWHEREAExpression)
	ctx.Step(`^UPDATE \'a(\d+)\' SET b = WHERE expression$`, uPDATEASETBWHEREExpression1)
	ctx.Step(`^UPDATE \'a(\d+)\' SET b WHERE expression$`, uPDATEASETBWHEREExpression2)
	ctx.Step(`^UPDATE \'a(\d+)\' SET expression$`, uPDATEASETExpression)
	ctx.Step(`^UPDATE expression$`, uPDATEExpression)
	ctx.Step(`^we get the query$`, weGetTheQuery)
}

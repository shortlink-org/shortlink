package v1

import (
	"fmt"
)

// Common errors -------------------------------------------------------------------------------------------------------
var ErrIncorrectSQLExpression = fmt.Errorf("incorrect sql-expression")
var ErrQueryTypeCannotBeEmpty = fmt.Errorf("query type cannot be empty")
var ErrTableNameCannotBeEmpty = fmt.Errorf("table name cannot be empty")

type ParserError struct {
	Err string
}

func (e *ParserError) Error() string {
	return e.Err
}

// Errors for SELECT ---------------------------------------------------------------------------------------------------
var ErrExpectedFieldToSelect = fmt.Errorf("at SELECT: expected field to SELECT")
var ErrExpectedCommaOrFrom = fmt.Errorf("at SELECT: expected comma or FROM")
var ErrExpectedFrom = fmt.Errorf("at SELECT: expected FROM")
var ErrExpectedQuotedTableName = fmt.Errorf("at SELECT: expected quoted table name")

type ExpectedFieldAliasToSelectError struct {
	Identifier string
}

func (e *ExpectedFieldAliasToSelectError) Error() string {
	return fmt.Sprintf("at SELECT: expected field alias for \"%s as\" to SELECT", e.Identifier)
}

// Errors for WHERE ----------------------------------------------------------------------------------------------------
var ErrExpectedField = fmt.Errorf("at WHERE: expected field")
var ErrExpectedOperator = fmt.Errorf("at WHERE: expected operator")
var ErrExpectedQuotedValue = fmt.Errorf("at WHERE: expected quoted value")
var ErrExpectedAnd = fmt.Errorf("at WHERE: expected AND")
var ErrEmptyWhereClause = fmt.Errorf("at WHERE: empty WHERE clause")
var ErrWhereClauseIsMandatory = fmt.Errorf("at WHERE: WHERE clause is mandatory for UPDATE & DELETE")
var ErrConditionWithoutOperator = fmt.Errorf("at WHERE: condition without operator")
var ErrConditionWithEmptyRightSideOperand = fmt.Errorf("at WHERE: condition with empty right side operand")
var ErrConditionWithEmptyLeftSideOperand = fmt.Errorf("at WHERE: condition with empty left side operand")

// Errors for INSERT INTO ----------------------------------------------------------------------------------------------
var ErrExpectedQuotedFieldName = fmt.Errorf("at INSERT INTO: expected quoted field name")
var ErrNeedAtLeastOneRowToInsert = fmt.Errorf("at INSERT INTO: need at least one row to insert")
var ErrValueCountDoesntMatchFieldCount = fmt.Errorf("at INSERT INTO: value count doesn't match field count")
var ErrExpectedQuotedFieldNameToInsert = fmt.Errorf("at INSERT INTO: expected quoted field name")
var ErrExpectedLessThanOneFieldToInsert = fmt.Errorf("at INSERT INTO: expected at least one field to insert")
var ErrExpectedValues = fmt.Errorf("at INSERT INTO: expected 'VALUES'")
var ErrExpectedOpeningParens = fmt.Errorf("at INSERT INTO: expected opening parens")
var ErrNotMatchedFieldAndValueCount = fmt.Errorf("at INSERT INTO: value count doesn't match field count")
var ErrExpectedCommaToInsert = fmt.Errorf("at INSERT INTO: expected comma")

// Errors for DELETE FROM ----------------------------------------------------------------------------------------------
var ErrExpectedQuotedTableNameToDelete = fmt.Errorf("at DELETE FROM: expected quoted table name")
var ErrExpectedWhere = fmt.Errorf("at DELETE FROM: expected WHERE")

// Errors for UPDATE ---------------------------------------------------------------------------------------------------
var ErrExpectedQuotedTableNameToUpdate = fmt.Errorf("at UPDATE: expected quoted table name")
var ErrExpectedSet = fmt.Errorf("at UPDATE: expected 'SET'")
var ErrExpectedQuotedFieldNameToUpdate = fmt.Errorf("at UPDATE: expected quoted field name to update")
var ErrEcpectedEqualSign = fmt.Errorf("at UPDATE: expected '='")
var ErrExpectedQuotedValueToUpdate = fmt.Errorf("at UPDATE: expected quoted value")
var ErrExpectedComma = fmt.Errorf("at UPDATE: expected ','")

// Errors for CREATE TABLE ---------------------------------------------------------------------------------------------
var ErrCreateTableTableNameCannotBeEmpty = fmt.Errorf("at CREATE TABLE: table name cannot be empty")
var ErrCreateTableExpectedOpeningParens = fmt.Errorf("at CREATE TABLE: expected opening parens")
var ErrCreateTableExpectedLessThanOneField = fmt.Errorf("at CREATE TABLE: expected at least one field to create table")
var ErrCreateTableExpectedQuotedFieldName = fmt.Errorf("at CREATE TABLE: expected quoted field name")
var ErrCreateTableUnsupportedTypeOfField = fmt.Errorf("at CREATE TABLE: unsupported type of field")
var ErrCreateTableExpectedCommaOrClosingParens = fmt.Errorf("at CREATE TABLE: expected comma or closing parens")

// Errors for LIMIT ----------------------------------------------------------------------------------------------------
var ErrEmptyLimitClause = fmt.Errorf("at LIMIT: empty LIMIT clause")
var ErrExpectedNumber = fmt.Errorf("at LIMIT: required number")

// Errors for JOIN -----------------------------------------------------------------------------------------------------
var ErrExpectedOperatorToJoin = fmt.Errorf("at ON: expected operator")
var ErrExpectedQuotedTableNameAndFieldNameToJoin = fmt.Errorf("at ON: expected <tablename>.<fieldname>")

// Errors for ORDER BY -------------------------------------------------------------------------------------------------
var ErrExpectedOrder = fmt.Errorf("expected ORDER")
var ErrExpectedFieldToOrder = fmt.Errorf("at ORDER BY: expected field to ORDER")

// Errors for INDEX ----------------------------------------------------------------------------------------------------
var ErrIncorrectSQLExpressionForIndex = fmt.Errorf("at INDEX: incorrect sql-expression")
var ErrExpectedQuotedIndexNameToDelete = fmt.Errorf("at DELETE INDEX: expected quoted index name")

// IncorrectTypeOfIndexError is an error for incorrect type of index
type IncorrectTypeOfIndexError struct {
	Type string
}

func (e *IncorrectTypeOfIndexError) Error() string {
	return fmt.Sprintf("at INDEX: incorrect type of index - %s", e.Type)
}

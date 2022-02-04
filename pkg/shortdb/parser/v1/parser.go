package v1

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/batazor/shortlink/internal/pkg/tool"
	v1 "github.com/batazor/shortlink/pkg/shortdb/query/v1"
	table "github.com/batazor/shortlink/pkg/shortdb/table/v1"
)

const (
	SELECT                  = "SELECT"
	INSERT_INTO             = "INSERT INTO"
	VALUES                  = "VALUES"
	UPDATE                  = "UPDATE"
	DELETE_FROM             = "DELETE FROM"
	WHERE                   = "WHERE"
	FROM                    = "FROM"
	SET                     = "SET"
	ON_DUPLICATE_KEY_UPDATE = "ON DUPLICATE KEY UPDATE"
	ORDER_BY                = "ORDER BY"
	ASC                     = "ASC"
	DESC                    = "DESC"
	LEFT_JOIN               = "LEFT JOIN"
	RIGHT_JOIN              = "RIGHT JOIN"
	INNER_JOIN              = "INNER_JOIN"
	JOIN                    = "JOIN"
	ON                      = "ON"
	AS                      = "AS"
	CREATE_TABLE            = "CREATE TABLE"
	DROP_TABLE              = "DROP_TABLE"
)

var reservedWords = []string{
	"(", ")", ">=", ">", "<", "<=", "=", "!=", ",", ";",
	SELECT, INSERT_INTO, VALUES, UPDATE,
	DELETE_FROM, WHERE, FROM, SET,
	ON_DUPLICATE_KEY_UPDATE, ORDER_BY, ASC,
	DESC, LEFT_JOIN, RIGHT_JOIN, INNER_JOIN,
	JOIN, ON, AS, CREATE_TABLE, DROP_TABLE,
}

var (
	r, _ = regexp.Compile("[a-zA-Z0-9]")

	// typeFieldTable - list of support type fields of table
	typeFieldTable = []string{"int", "integer", "string", "text", "boolean", "bool"}
)

func New(sql string) (*Parser, error) {
	parser := &Parser{
		Sql:   sql,
		Query: &v1.Query{},
	}

	// Parse
	_, err := parser.Parse()
	if err != nil {
		return parser, err
	}

	return parser, nil
}

// Parse - main function that returns the "query struct" or an error
func (p *Parser) Parse() (*v1.Query, error) {
	q, err := p.doParse()
	p.Error = err.Error()
	if p.Error == "" {
		if errValidate := p.validate(); errValidate != nil {
			p.Error = errValidate.Error()
		}
	}
	return q, fmt.Errorf(p.Error)
}

func (p *Parser) doParse() (*v1.Query, error) { // nolint gocyclo
	for {
		if p.I >= int32(len(p.Sql)) {
			return p.Query, fmt.Errorf(p.Error)
		}

		switch p.Step {
		case Step_STEP_UNSPECIFIED:
			switch strings.ToUpper(p.peek()) {
			case "SELECT":
				p.Query.Type = v1.Type_TYPE_SELECT
				p.pop()
				p.Step = Step_STEP_SELECT_FIELD
			case "INSERT INTO":
				p.Query.Type = v1.Type_TYPE_INSERT
				p.pop()
				p.Step = Step_STEP_INSERT_TABLE
			case "UPDATE":
				p.Query.Type = v1.Type_TYPE_UPDATE
				p.Query.Updates = map[string]string{}
				p.pop()
				p.Step = Step_STEP_UPDATE_TABLE
			case "DELETE FROM":
				p.Query.Type = v1.Type_TYPE_DELETE
				p.pop()
				p.Step = Step_STEP_UPDATE_TABLE
			case "CREATE TABLE":
				p.Query.Type = v1.Type_TYPE_CREATE_TABLE
				p.pop()
				p.Step = Step_STEP_CREATE_TABLE_NAME
			case "DROP TABLE":
				p.Query.Type = v1.Type_TYPE_DROP_TABLE
				p.pop()
				p.Step = Step_STEP_DROP_TABLE_NAME
			default:
				return nil, fmt.Errorf("incorrect sql-expression")
			}
		case Step_STEP_SEMICOLON:
			p.pop()
		case Step_STEP_SELECT_FIELD:
			identifier := p.peek()
			if !isIdentifierOrAsterisk(identifier) {
				return p.Query, fmt.Errorf("at SELECT: expected field to SELECT")
			}

			p.Query.Fields = append(p.Query.Fields, identifier)
			p.pop()
			maybeFrom := p.peek()

			if strings.ToUpper(maybeFrom) == "AS" {
				p.pop()
				alias := p.peek()
				if !isIdentifier(alias) {
					return p.Query, fmt.Errorf("at SELECT: expected field alias for \"%s as\" to SELECT", identifier)
				}
				if p.Query.Aliases == nil {
					p.Query.Aliases = make(map[string]string)
				}
				p.Query.Aliases[identifier] = alias
				p.pop()
				maybeFrom = p.peek()
			}

			if strings.ToUpper(maybeFrom) == "FROM" {
				p.Step = Step_STEP_SELECT_FROM
				continue
			}

			p.Step = Step_STEP_SELECT_COMMA
		case Step_STEP_SELECT_COMMA:
			commaRWord := p.peek()
			if commaRWord != "," {
				return p.Query, fmt.Errorf("at SELECT: expected comma or FROM")
			}
			p.pop()
			p.Step = Step_STEP_SELECT_FIELD
		case Step_STEP_SELECT_FROM:
			fromRWord := p.peek()
			if strings.ToUpper(fromRWord) != "FROM" {
				return p.Query, fmt.Errorf("at SELECT: expected FROM")
			}
			p.pop()
			p.Step = Step_STEP_SELECT_FROM_TABLE
		case Step_STEP_SELECT_FROM_TABLE:
			tableName := p.peek()
			if len(tableName) == 0 {
				return p.Query, fmt.Errorf("at SELECT: expected quoted table name")
			}

			if strings.Contains(tableName, ".") {
				parts := strings.Split(tableName, ".")
				p.Query.Database = parts[0]
				tableName = parts[1]
			}

			p.Query.TableName = tableName
			p.pop()
			look := p.peek()
			if strings.ToUpper(look) == WHERE {
				p.Step = Step_STEP_WHERE
			} else if strings.ToUpper(look) == ORDER_BY {
				p.Step = Step_STEP_ORDER
			} else if strings.Contains(strings.ToUpper(look), JOIN) {
				p.Step = Step_STEP_JOIN
			} else if look == ";" {
				p.Step = Step_STEP_SEMICOLON
			}
		case Step_STEP_INSERT_TABLE:
			tableName := p.peek()
			if len(tableName) == 0 {
				return p.Query, fmt.Errorf("at INSERT INTO: expected quoted table name")
			}
			p.Query.TableName = tableName
			p.pop()
			p.Step = Step_STEP_INSERT_FIELD_OPENING_PARENTS
		case Step_STEP_DELETE_FROM_TABLE:
			tableName := p.peek()
			if len(tableName) == 0 {
				return p.Query, fmt.Errorf("at DELETE FROM: expected quoted table name")
			}
			p.Query.TableName = tableName
			p.pop()
			p.Step = Step_STEP_WHERE
		case Step_STEP_WHERE:
			whereRWord := p.peek()
			if strings.ToUpper(whereRWord) != WHERE {
				return p.Query, fmt.Errorf("expected WHERE")
			}
			p.pop()
			p.Step = Step_STEP_WHERE_FIELD
		case Step_STEP_WHERE_FIELD:
			identifier := p.peek()
			if !isIdentifier(identifier) {
				return p.Query, fmt.Errorf("at WHERE: expected field")
			}
			p.Query.Conditions = append(p.Query.Conditions, &v1.Condition{LValue: identifier, LValueIsField: true})
			p.pop()
			p.Step = Step_STEP_WHERE_OPERATOR
		case Step_STEP_WHERE_OPERATOR:
			currentCondition := p.Query.Conditions[len(p.Query.Conditions)-1]

			operator := p.peek()
			currentCondition.Operator = getOperator(operator)
			if currentCondition.Operator == v1.Operator_OPERATOR_UNSPECIFIED {
				return p.Query, fmt.Errorf("at WHERE: unknown operator")
			}

			p.Query.Conditions[len(p.Query.Conditions)-1] = currentCondition
			p.pop()
			p.Step = Step_STEP_WHERE_VALUE
		case Step_STEP_WHERE_VALUE:
			currentCondition := p.Query.Conditions[len(p.Query.Conditions)-1]
			identifier := p.peek()
			if isIdentifier(identifier) {
				currentCondition.RValue = identifier
				currentCondition.RValueIsField = true
			} else {
				quotedValue, ln := p.peekQuotedStringWithLength()
				if ln == 0 {
					return p.Query, fmt.Errorf("at WHERE: expected quoted value")
				}
				currentCondition.RValue = quotedValue
				currentCondition.RValueIsField = false
			}
			p.Query.Conditions[len(p.Query.Conditions)-1] = currentCondition
			p.pop()
			p.Step = Step_STEP_WHERE_AND
		case Step_STEP_WHERE_AND:
			andRWord := p.peek()
			if strings.ToUpper(andRWord) != "AND" {
				return p.Query, fmt.Errorf("expected AND")
			}
			p.pop()
			p.Step = Step_STEP_WHERE_FIELD
		case Step_STEP_UPDATE_TABLE:
			tableName := p.peek()
			if len(tableName) == 0 {
				return p.Query, fmt.Errorf("at UPDATE: expected quoted table name")
			}
			p.Query.TableName = tableName
			p.pop()
			p.Step = Step_STEP_UPDATE_SET
		case Step_STEP_UPDATE_SET:
			setRWord := p.peek()
			if setRWord != "SET" {
				return p.Query, fmt.Errorf("at UPDATE: expected 'SET'")
			}
			p.pop()
			p.Step = Step_STEP_UPDATE_FIELD
		case Step_STEP_UPDATE_FIELD:
			identifier := p.peek()
			if !isIdentifier(identifier) {
				return p.Query, fmt.Errorf("at UPDATE: expected at least one field to update")
			}
			p.NextUpdateField = identifier
			p.pop()
			p.Step = Step_STEP_UPDATE_EQUALS
		case Step_STEP_UPDATE_EQUALS:
			equalsRWord := p.peek()
			if equalsRWord != "=" {
				return p.Query, fmt.Errorf("at UPDATE: expected '='")
			}
			p.pop()
			p.Step = Step_STEP_UPDATE_VALUE
		case Step_STEP_UPDATE_VALUE:
			quotedValue, ln := p.peekQuotedStringWithLength()
			if ln == 0 {
				return p.Query, fmt.Errorf("at UPDATE: expected quoted value")
			}
			p.Query.Updates[p.NextUpdateField] = quotedValue
			p.NextUpdateField = ""
			p.pop()
			maybeWhere := p.peek()
			if strings.ToUpper(maybeWhere) == "WHERE" {
				p.Step = Step_STEP_WHERE
				continue
			}
			p.Step = Step_STEP_UPDATE_COMMA
		case Step_STEP_UPDATE_COMMA:
			commaRWord := p.peek()
			if commaRWord != "," {
				return p.Query, fmt.Errorf("at UPDATE: expected ','")
			}
			p.pop()
			p.Step = Step_STEP_UPDATE_FIELD
		case Step_STEP_ORDER:
			orderRWord := p.peek()
			if strings.ToUpper(orderRWord) != "ORDER BY" {
				return p.Query, fmt.Errorf("expected ORDER")
			}
			p.pop()
			p.Step = Step_STEP_ORDER_FIELD
		case Step_STEP_ORDER_FIELD:
			identifier := p.peek()
			if !isIdentifier(identifier) {
				return p.Query, fmt.Errorf("at ORDER BY: expected field to ORDER")
			}
			p.Query.OrderFields = append(p.Query.OrderFields, identifier)
			p.Query.OrderDir = append(p.Query.OrderDir, "ASC")
			p.pop()
			p.Step = Step_STEP_ORDER_DIRECTION_OR_COMMA
		case Step_STEP_ORDER_DIRECTION_OR_COMMA:
			commaRWord := p.peek()
			if commaRWord == "," {
				p.pop()
			} else if commaRWord == "ASC" || commaRWord == "DESC" {
				p.pop()
				p.Query.OrderDir[len(p.Query.OrderDir)-1] = commaRWord
				continue
			}
			p.Step = Step_STEP_ORDER_FIELD
		case Step_STEP_JOIN:
			joinType := p.peek()
			p.Query.Joins = append(p.Query.Joins, &v1.Join{Type: joinType, Table: "UNKNOWN"})
			p.pop()
			p.Step = Step_STEP_JOIN_TABLE
		case Step_STEP_JOIN_TABLE:
			joinTable := p.peek()
			currentJoin := p.Query.Joins[len(p.Query.Joins)-1]
			currentJoin.Table = joinTable
			p.Query.Joins[len(p.Query.Joins)-1] = currentJoin
			p.pop()
			if strings.ToUpper(p.peek()) == "ON" {
				p.Step = Step_STEP_JOIN_CONDITION
			} else {
				p.Step = Step_STEP_ORDER
			}
		case Step_STEP_JOIN_CONDITION:
			p.pop()
			op1 := p.pop()
			op1split := strings.Split(op1, ".")
			if len(op1split) != 2 {
				return p.Query, fmt.Errorf("at ON: expected <tablename>.<fieldname>")
			}
			currentCondition := &v1.JoinCondition{LTable: op1split[0], LOperand: op1split[1]}

			operator := p.peek()
			currentCondition.Operator = getOperator(operator)
			if currentCondition.Operator == v1.Operator_OPERATOR_UNSPECIFIED {
				return p.Query, fmt.Errorf("at ON: unknown operator")
			}

			p.pop()
			op2 := p.pop()
			op2split := strings.Split(op2, ".")
			if len(op2split) != 2 {
				return p.Query, fmt.Errorf("at ON: expected <tablename>.<fieldname>")
			}
			currentCondition.RTable = op2split[0]
			currentCondition.ROperand = op2split[1]
			currentJoin := p.Query.Joins[len(p.Query.Joins)-1]
			currentJoin.Conditions = append(currentJoin.Conditions, currentCondition)
			p.Query.Joins[len(p.Query.Joins)-1] = currentJoin
			nextOp := p.peek()
			if strings.ToUpper(nextOp) == "WHERE" {
				p.Step = Step_STEP_WHERE
			} else if strings.ToUpper(nextOp) == "ORDER BY" {
				p.Step = Step_STEP_ORDER
			} else if strings.ToUpper(nextOp) == "AND" {
				p.Step = Step_STEP_JOIN_CONDITION
			} else if strings.Contains(strings.ToUpper(nextOp), "JOIN") {
				p.Step = Step_STEP_JOIN
			}
		case Step_STEP_INSERT_FIELD_OPENING_PARENTS:
			openingParens := p.peek()
			if len(openingParens) != 1 || openingParens != "(" {
				return p.Query, fmt.Errorf("at INSERT INTO: expected opening parens")
			}
			p.pop()
			p.Step = Step_STEP_INSERT_FIELDS
		case Step_STEP_INSERT_FIELDS:
			identifier := p.peek()
			if !isIdentifier(identifier) {
				return p.Query, fmt.Errorf("at INSERT INTO: expected at least one field to insert")
			}
			p.Query.Fields = append(p.Query.Fields, identifier)
			p.pop()
			p.Step = Step_STEP_INSERT_FIELDS_COMMA_OR_CLOSING_PARENTS
		case Step_STEP_INSERT_FIELDS_COMMA_OR_CLOSING_PARENTS:
			commaOrClosingParens := p.peek()
			if commaOrClosingParens != "," && commaOrClosingParens != ")" {
				return p.Query, fmt.Errorf("at INSERT INTO: expected comma or closing parens")
			}
			p.pop()
			if commaOrClosingParens == "," {
				p.Step = Step_STEP_INSERT_FIELDS
				continue
			}
			p.Step = Step_STEP_INSERT_RWORD
		case Step_STEP_INSERT_RWORD:
			valuesRWord := p.peek()
			if strings.ToUpper(valuesRWord) != "VALUES" {
				return p.Query, fmt.Errorf("at INSERT INTO: expected 'VALUES'")
			}
			p.pop()
			p.Step = Step_STEP_INSERT_VALUES_OPENING_PARENS
		case Step_STEP_INSERT_VALUES_OPENING_PARENS:
			openingParens := p.peek()
			if openingParens != "(" {
				return p.Query, fmt.Errorf("at INSERT INTO: expected opening parens")
			}
			p.Query.Inserts = append(p.Query.Inserts, &v1.Query_Array{})
			p.pop()
			p.Step = Step_STEP_INSERT_VALUES
		case Step_STEP_INSERT_VALUES:
			quotedValue, ln := p.peekQuotedStringWithLength()
			if ln == 0 {
				return p.Query, fmt.Errorf("at INSERT INTO: expected quoted value")
			}
			p.Query.Inserts[len(p.Query.Inserts)-1].Items = append(p.Query.Inserts[len(p.Query.Inserts)-1].Items, quotedValue)
			p.pop()
			p.Step = Step_STEP_INSERT_VALUES_COMMA_OR_CLOSING_PARENS
		case Step_STEP_INSERT_VALUES_COMMA_OR_CLOSING_PARENS:
			commaOrClosingParens := p.peek()
			if commaOrClosingParens != "," && commaOrClosingParens != ")" {
				return p.Query, fmt.Errorf("at INSERT INTO: expected comma or closing parens")
			}
			p.pop()
			if commaOrClosingParens == "," {
				p.Step = Step_STEP_INSERT_VALUES
				continue
			}
			currentInsertRow := p.Query.Inserts[len(p.Query.Inserts)-1]
			if len(currentInsertRow.Items) < len(p.Query.Fields) {
				return p.Query, fmt.Errorf("at INSERT INTO: value count doesn't match field count")
			}
			p.Step = Step_STEP_INSERT_VALUES_COMMA_BEFORE_OPENING_PARENS
		case Step_STEP_INSERT_VALUES_COMMA_BEFORE_OPENING_PARENS:
			commaRWord := p.peek()
			if commaRWord == ";" {
				p.Step = Step_STEP_SEMICOLON
			} else if strings.ToUpper(commaRWord) != "," {
				return p.Query, fmt.Errorf("at INSERT INTO: expected comma")
			}
			p.pop()
			p.Step = Step_STEP_INSERT_VALUES_OPENING_PARENS
		case Step_STEP_CREATE_TABLE_NAME:
			identifier := p.peek()
			if !isIdentifier(identifier) {
				return p.Query, fmt.Errorf("at CREATE TABLE: table name cannot be empty")
			}
			p.pop()
			p.Query.TableName = identifier
			p.Step = Step_STEP_CREATE_TABLE_OPENING_PARENS
		case Step_STEP_CREATE_TABLE_OPENING_PARENS:
			openingParens := p.peek()
			p.pop()
			if openingParens != "(" {
				return p.Query, fmt.Errorf("at CREATE TABLE: expected opening parens")
			}

			p.Step = Step_STEP_CREATE_TABLE_FIELDS
		case Step_STEP_CREATE_TABLE_FIELDS:
			// get name field of table
			identifier := p.peek()
			if identifier == ")" {
				p.Step = Step_STEP_CREATE_TABLE_FIELDS_COMMA_OR_CLOSING_PARENS
				continue
			}
			if !isIdentifier(identifier) {
				return p.Query, fmt.Errorf("at CREATE TABLE: expected at least one field to create table")
			}

			p.pop()
			typeField := p.peek()

			// get type field of table
			if !isIdentifier(typeField) {
				return p.Query, fmt.Errorf("at CREATE TABLE: expected at least one field to create table")
			}

			if p.Query.TableFields == nil {
				p.Query.TableFields = map[string]table.Type{}
			}

			// append field to table
			if tool.Contains(typeFieldTable, typeField) {
				switch typeField {
				case "int":
					fallthrough
				case "integer":
					p.Query.TableFields[identifier] = table.Type_TYPE_INTEGER
				case "text":
					fallthrough
				case "string":
					p.Query.TableFields[identifier] = table.Type_TYPE_STRING
				case "bool":
					fallthrough
				case "boolean":
					p.Query.TableFields[identifier] = table.Type_TYPE_BOOLEAN
				default:
					return p.Query, fmt.Errorf("at CREATE TABLE: unsupported type of field")
				}

				p.pop()
			} else {
				return p.Query, fmt.Errorf("at CREATE TABLE: unsupported type of field")
			}

			p.Step = Step_STEP_CREATE_TABLE_FIELDS_COMMA_OR_CLOSING_PARENS
		case Step_STEP_CREATE_TABLE_FIELDS_COMMA_OR_CLOSING_PARENS:
			commaOrClosingParens := p.peek()
			if commaOrClosingParens != "," && commaOrClosingParens != ")" && commaOrClosingParens != ";" {
				return p.Query, fmt.Errorf("at CREATE TABLE: expected comma or closing parens")
			}
			p.pop()
			if commaOrClosingParens == "," {
				p.Step = Step_STEP_CREATE_TABLE_FIELDS
				continue
			}
			if commaOrClosingParens == ";" {
				p.Step = Step_STEP_SEMICOLON
			}
		case Step_STEP_DROP_TABLE_NAME:
			commaRWord := p.peek()
			p.pop()
			p.Query.TableName = commaRWord
		}
	}
}

// peek - a "look-ahead" function that returns the next token to parse
func (p *Parser) peek() string {
	peeked, _ := p.peekWithLength()
	return peeked
}

// pop - same as peek(), but advancing our "i" index
func (p *Parser) pop() string {
	peeked, len := p.peekWithLength()
	p.I += len
	p.popWhitespace()
	return peeked
}

func (p *Parser) peekWithLength() (string, int32) {
	if p.I >= int32(len(p.Sql)) {
		return "", 0
	}

	for _, rWord := range reservedWords {
		token := strings.ToUpper(p.Sql[p.I:min(len(p.Sql), int(p.I)+len(rWord))])
		if token == rWord {
			return token, int32(len(token))
		}
	}

	if p.Sql[p.I] == '\'' { // Quoted string
		return p.peekQuotedStringWithLength()
	}

	return p.peekIdentifierStringWithLength()
}

func (p *Parser) popWhitespace() {
	for ; p.I < int32(len(p.Sql)) && p.Sql[p.I] == ' '; p.I++ {
	}
}

func (p *Parser) peekQuotedStringWithLength() (string, int32) {
	if int32(len(p.Sql)) < p.I || p.Sql[p.I] != '\'' {
		return "", 0
	}

	for i := p.I + 1; i < int32(len(p.Sql)); i++ {
		if p.Sql[i] == '\'' && p.Sql[i-1] != '\\' {
			return p.Sql[p.I+1 : i], int32(len(p.Sql[p.I+1:i]) + 2) // +2 for the two quotes
		}
	}

	return "", 0
}

func (p *Parser) peekIdentifierStringWithLength() (string, int32) {
	for i := p.I; i < int32(len(p.Sql)); i++ {
		if matched := r.MatchString(string(p.Sql[i])); !matched {
			return p.Sql[p.I:i], int32(len(p.Sql[p.I:i]))
		}
	}

	return p.Sql[p.I:], int32(len(p.Sql[p.I:]))
}

func (p *Parser) validate() error { // nolint gocyclo
	if p.Query == nil {
		return nil
	}

	if len(p.Query.Conditions) == 0 && p.Step == Step_STEP_WHERE_FIELD {
		return fmt.Errorf("at WHERE: empty WHERE clause")
	}

	if p.Query.Type == v1.Type_TYPE_UNSPECIFIED {
		return fmt.Errorf("query type cannot be empty")
	}

	if p.Query.TableName == "" {
		return fmt.Errorf("table name cannot be empty")
	}

	if len(p.Query.Conditions) == 0 && (p.Query.Type == v1.Type_TYPE_UPDATE || p.Query.Type == v1.Type_TYPE_DELETE) {
		return fmt.Errorf("at WHERE: WHERE clause is mandatory for UPDATE & DELETE")
	}

	for _, c := range p.Query.Conditions {
		if c.Operator == v1.Operator_OPERATOR_UNSPECIFIED {
			return fmt.Errorf("at WHERE: condition without operator")
		}

		if c.LValue == "" && c.LValueIsField {
			return fmt.Errorf("at WHERE: condition with empty left side operand")
		}

		if c.RValue == "" && c.RValueIsField {
			return fmt.Errorf("at WHERE: condition with empty right side operand")
		}
	}

	if p.Query.Type == v1.Type_TYPE_INSERT && len(p.Query.Inserts) == 0 {
		return fmt.Errorf("at INSERT INTO: need at least one row to insert")
	}

	if p.Query.Type == v1.Type_TYPE_INSERT {
		for _, i := range p.Query.Inserts {
			if len(i.Items) != len(p.Query.Fields) {
				return fmt.Errorf("at INSERT INTO: value count doesn't match field count")
			}
		}
	}

	return nil
}

func isIdentifier(s string) bool {
	for _, rw := range reservedWords {
		if strings.ToUpper(s) == rw {
			return false
		}
	}

	matched, _ := regexp.MatchString("[a-zA-Z_][a-zA-Z_0-9]*", s)
	return matched
}

func isIdentifierOrAsterisk(s string) bool {
	return isIdentifier(s) || s == "*"
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getOperator(operator string) v1.Operator {
	switch operator {
	case "=":
		return v1.Operator_OPERATOR_EQ
	case ">":
		return v1.Operator_OPERATOR_GT
	case ">=":
		return v1.Operator_OPERATOR_GTE
	case "<":
		return v1.Operator_OPERATOR_LT
	case "<=":
		return v1.Operator_OPERATOR_LTE
	case "!=":
		return v1.Operator_OPERATOR_NE
	default:
		return v1.Operator_OPERATOR_UNSPECIFIED
	}
}

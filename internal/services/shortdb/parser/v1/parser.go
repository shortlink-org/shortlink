package v1

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/shortlink-org/shortlink/internal/pkg/types/vector"
	field "github.com/shortlink-org/shortlink/internal/services/shortdb/domain/field/v1"
	"github.com/shortlink-org/shortlink/internal/services/shortdb/domain/index/v1"
	query "github.com/shortlink-org/shortlink/internal/services/shortdb/domain/query/v1"
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
	LIMIT                   = "LIMIT"
	CREATE_INDEX            = "CREATE INDEX"
	DELETE_INDEX            = "DELETE INDEX"
)

var reservedWords = []string{
	"(", ")", ">=", ">", "<", "<=", "=", "!=", ",", ";",
	SELECT, INSERT_INTO, VALUES, UPDATE,
	DELETE_FROM, WHERE, FROM, SET,
	ON_DUPLICATE_KEY_UPDATE, ORDER_BY, ASC,
	DESC, LEFT_JOIN, RIGHT_JOIN, INNER_JOIN,
	JOIN, ON, AS, CREATE_TABLE, DROP_TABLE,
	LIMIT, CREATE_INDEX, DELETE_INDEX,
}

var (
	r = regexp.MustCompile("[a-zA-Z0-9]") //nolint:errcheck // TODO: refactor

	// TypeFieldTable - list of support type fields of table
	typeFieldTable = []string{"int", "integer", "string", "text", "boolean", "bool"}
)

func New(sql string) (*Parser, error) {
	parser := &Parser{
		Sql:   sql,
		Query: query.New(),
	}

	// Parse
	_, err := parser.Parse()
	if err != nil {
		return parser, err
	}

	return parser, nil
}

// Parse - main function that returns the "query struct" or an error
func (p *Parser) Parse() (*query.Query, error) {
	q, err := p.doParse()
	p.Error = err.Error()

	if p.GetError() == "" {
		if errValidate := p.validate(); errValidate != nil {
			p.Error = errValidate.Error()
		}
	}

	if p.GetError() != "" {
		return nil, fmt.Errorf(p.Error)
	}

	return q, nil
}

func (p *Parser) doParse() (*query.Query, error) { //nolint:gocyclo,gocognit,maintidx,revive,cyclop // TODO: refactor
	for {
		if p.GetI() >= int32(len(p.GetSql())) {
			return p.Query, fmt.Errorf(p.Error)
		}

		switch p.GetStep() {
		case Step_STEP_UNSPECIFIED:
			switch strings.ToUpper(p.peek()) {
			case "SELECT":
				p.Query.Type = query.Type_TYPE_SELECT
				p.pop()
				p.Step = Step_STEP_SELECT_FIELD
			case "INSERT INTO":
				p.Query.Type = query.Type_TYPE_INSERT
				p.pop()
				p.Step = Step_STEP_INSERT_TABLE
			case "UPDATE":
				p.Query.Type = query.Type_TYPE_UPDATE
				p.Query.Updates = map[string]string{}
				p.pop()
				p.Step = Step_STEP_UPDATE_TABLE
			case "DELETE FROM":
				p.Query.Type = query.Type_TYPE_DELETE
				p.pop()
				p.Step = Step_STEP_UPDATE_TABLE
			case "CREATE TABLE":
				p.Query.Type = query.Type_TYPE_CREATE_TABLE
				p.pop()
				p.Step = Step_STEP_CREATE_TABLE_NAME
			case "DROP TABLE":
				p.Query.Type = query.Type_TYPE_DROP_TABLE
				p.pop()
				p.Step = Step_STEP_DROP_TABLE_NAME
			case "CREATE INDEX":
				p.Query.Type = query.Type_TYPE_CREATE_INDEX
				p.pop()
				p.Step = Step_STEP_CREATE_INDEX_NAME
			default:
				return nil, fmt.Errorf("incorrect sql-expression")
			}
		case Step_STEP_SEMICOLON:
			p.pop()
		case Step_STEP_SELECT_FIELD:
			identifier := p.peek()
			if !isIdentifierOrAsterisk(identifier) {
				return p.GetQuery(), fmt.Errorf("at SELECT: expected field to SELECT")
			}

			p.Query.Fields = append(p.GetQuery().GetFields(), identifier)
			p.pop()
			maybeFrom := p.peek()

			if strings.EqualFold(maybeFrom, "AS") {
				p.pop()

				alias := p.peek()
				if !isIdentifier(alias) {
					return p.GetQuery(), fmt.Errorf("at SELECT: expected field alias for \"%s as\" to SELECT", identifier)
				}

				if p.GetQuery().GetAliases() == nil {
					p.Query.Aliases = make(map[string]string)
				}

				p.Query.Aliases[identifier] = alias
				p.pop()
				maybeFrom = p.peek()
			}

			if strings.EqualFold(maybeFrom, "FROM") {
				p.Step = Step_STEP_SELECT_FROM
				continue
			}

			p.Step = Step_STEP_SELECT_COMMA
		case Step_STEP_SELECT_COMMA:
			commaRWord := p.peek()
			if commaRWord != "," {
				return p.GetQuery(), fmt.Errorf("at SELECT: expected comma or FROM")
			}

			p.pop()
			p.Step = Step_STEP_SELECT_FIELD
		case Step_STEP_SELECT_FROM:
			fromRWord := p.peek()
			if !strings.EqualFold(fromRWord, "FROM") {
				return p.GetQuery(), fmt.Errorf("at SELECT: expected FROM")
			}

			p.pop()
			p.Step = Step_STEP_SELECT_FROM_TABLE
		case Step_STEP_SELECT_FROM_TABLE:
			tableName := p.peek()
			if tableName == "" {
				return p.GetQuery(), fmt.Errorf("at SELECT: expected quoted table name")
			}

			if strings.Contains(tableName, ".") {
				parts := strings.Split(tableName, ".")
				p.Query.Database = parts[0]
				tableName = parts[1]
			}

			p.Query.TableName = tableName
			p.pop()

			look := p.peek()
			if strings.EqualFold(look, WHERE) {
				p.Step = Step_STEP_WHERE
			} else if strings.EqualFold(look, ORDER_BY) {
				p.Step = Step_STEP_ORDER
			} else if strings.Contains(strings.ToUpper(look), JOIN) {
				p.Step = Step_STEP_JOIN
			} else if strings.EqualFold(look, LIMIT) {
				p.Step = Step_STEP_LIMIT
			} else if look == ";" {
				p.Step = Step_STEP_SEMICOLON
			}
		case Step_STEP_INSERT_TABLE:
			tableName := p.peek()
			if tableName == "" {
				return p.GetQuery(), fmt.Errorf("at INSERT INTO: expected quoted table name")
			}

			p.Query.TableName = tableName
			p.pop()
			p.Step = Step_STEP_INSERT_FIELD_OPENING_PARENTS
		case Step_STEP_DELETE_FROM_TABLE:
			tableName := p.peek()
			if tableName == "" {
				return p.GetQuery(), fmt.Errorf("at DELETE FROM: expected quoted table name")
			}

			p.Query.TableName = tableName
			p.pop()
			p.Step = Step_STEP_WHERE
		case Step_STEP_WHERE:
			whereRWord := p.peek()
			if !strings.EqualFold(whereRWord, WHERE) {
				return p.GetQuery(), fmt.Errorf("expected WHERE")
			}
			p.pop()
			p.Step = Step_STEP_WHERE_FIELD
		case Step_STEP_WHERE_FIELD:
			identifier := p.peek()
			if !isIdentifier(identifier) {
				return p.GetQuery(), fmt.Errorf("at WHERE: expected field")
			}

			p.Query.Conditions = append(p.GetQuery().GetConditions(), &query.Condition{LValue: identifier, LValueIsField: true})
			p.pop()
			p.Step = Step_STEP_WHERE_OPERATOR
		case Step_STEP_WHERE_OPERATOR:
			currentCondition := p.GetQuery().GetConditions()[len(p.GetQuery().GetConditions())-1]

			operator := p.peek()
			currentCondition.Operator = getOperator(operator)
			if currentCondition.GetOperator() == query.Operator_OPERATOR_UNSPECIFIED {
				return p.GetQuery(), fmt.Errorf("at WHERE: unknown operator")
			}

			p.Query.Conditions[len(p.GetQuery().GetConditions())-1] = currentCondition
			p.pop()
			p.Step = Step_STEP_WHERE_VALUE
		case Step_STEP_WHERE_VALUE:
			currentCondition := p.GetQuery().GetConditions()[len(p.GetQuery().GetConditions())-1]

			identifier := p.peek()
			if isIdentifier(identifier) {
				currentCondition.RValue = identifier
				currentCondition.RValueIsField = true
			} else {
				quotedValue, ln := p.peekQuotedStringWithLength()
				if ln == 0 {
					return p.GetQuery(), fmt.Errorf("at WHERE: expected quoted value")
				}
				currentCondition.RValue = quotedValue
				currentCondition.RValueIsField = false
			}
			p.Query.Conditions[len(p.GetQuery().GetConditions())-1] = currentCondition
			p.pop()

			operator := p.peek()
			if strings.EqualFold(operator, LIMIT) {
				p.Step = Step_STEP_LIMIT
				continue
			}

			p.Step = Step_STEP_WHERE_AND
		case Step_STEP_WHERE_AND:
			andRWord := p.peek()
			if !strings.EqualFold(andRWord, "AND") {
				return p.GetQuery(), fmt.Errorf("expected AND")
			}

			p.pop()
			p.Step = Step_STEP_WHERE_FIELD
		case Step_STEP_UPDATE_TABLE:
			tableName := p.peek()
			if tableName == "" {
				return p.GetQuery(), fmt.Errorf("at UPDATE: expected quoted table name")
			}

			p.Query.TableName = tableName
			p.pop()
			p.Step = Step_STEP_UPDATE_SET
		case Step_STEP_UPDATE_SET:
			setRWord := p.peek()
			if setRWord != "SET" {
				return p.GetQuery(), fmt.Errorf("at UPDATE: expected 'SET'")
			}

			p.pop()
			p.Step = Step_STEP_UPDATE_FIELD
		case Step_STEP_UPDATE_FIELD:
			identifier := p.peek()
			if !isIdentifier(identifier) {
				return p.GetQuery(), fmt.Errorf("at UPDATE: expected at least one field to update")
			}

			p.NextUpdateField = identifier
			p.pop()
			p.Step = Step_STEP_UPDATE_EQUALS
		case Step_STEP_UPDATE_EQUALS:
			equalsRWord := p.peek()
			if equalsRWord != "=" {
				return p.GetQuery(), fmt.Errorf("at UPDATE: expected '='")
			}

			p.pop()
			p.Step = Step_STEP_UPDATE_VALUE
		case Step_STEP_UPDATE_VALUE:
			quotedValue, ln := p.peekQuotedStringWithLength()
			if ln == 0 {
				return p.GetQuery(), fmt.Errorf("at UPDATE: expected quoted value")
			}

			p.Query.Updates[p.GetNextUpdateField()] = quotedValue
			p.NextUpdateField = ""
			p.pop()

			maybeWhere := p.peek()
			if strings.EqualFold(maybeWhere, "WHERE") {
				p.Step = Step_STEP_WHERE
				continue
			}

			p.Step = Step_STEP_UPDATE_COMMA
		case Step_STEP_UPDATE_COMMA:
			commaRWord := p.peek()
			if commaRWord != "," {
				return p.GetQuery(), fmt.Errorf("at UPDATE: expected ','")
			}

			p.pop()
			p.Step = Step_STEP_UPDATE_FIELD
		case Step_STEP_DELETE_INDEX:
			indexName := p.peek()
			if indexName == "" {
				return p.GetQuery(), fmt.Errorf("at DELETE INDEX: expected quoted index name")
			}

			p.pop()
		case Step_STEP_ORDER:
			orderRWord := p.peek()
			if !strings.EqualFold(orderRWord, "ORDER BY") {
				return p.GetQuery(), fmt.Errorf("expected ORDER")
			}

			p.pop()
			p.Step = Step_STEP_ORDER_FIELD
		case Step_STEP_ORDER_FIELD:
			identifier := p.peek()
			if !isIdentifier(identifier) {
				return p.GetQuery(), fmt.Errorf("at ORDER BY: expected field to ORDER")
			}

			p.Query.OrderFields = append(p.GetQuery().GetOrderFields(), identifier)
			p.Query.OrderDir = append(p.GetQuery().GetOrderDir(), "ASC")
			p.pop()
			p.Step = Step_STEP_ORDER_DIRECTION_OR_COMMA
		case Step_STEP_ORDER_DIRECTION_OR_COMMA:
			commaRWord := p.peek()
			if commaRWord == "," {
				p.pop()
			} else if commaRWord == "ASC" || commaRWord == "DESC" {
				p.pop()
				p.Query.OrderDir[len(p.GetQuery().GetOrderDir())-1] = commaRWord

				continue
			}

			p.Step = Step_STEP_ORDER_FIELD
		case Step_STEP_JOIN:
			joinType := p.peek()
			p.Query.Joins = append(p.GetQuery().GetJoins(), &query.Join{Type: joinType, Table: "UNKNOWN"})
			p.pop()
			p.Step = Step_STEP_JOIN_TABLE
		case Step_STEP_JOIN_TABLE:
			joinTable := p.peek()
			currentJoin := p.GetQuery().GetJoins()[len(p.GetQuery().GetJoins())-1]
			currentJoin.Table = joinTable
			p.Query.Joins[len(p.GetQuery().GetJoins())-1] = currentJoin
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
			if len(op1split) != 2 { //nolint:gomnd // ignore
				return p.GetQuery(), fmt.Errorf("at ON: expected <tablename>.<fieldname>")
			}
			currentCondition := &query.JoinCondition{LTable: op1split[0], LOperand: op1split[1]}

			operator := p.peek()
			currentCondition.Operator = getOperator(operator)
			if currentCondition.GetOperator() == query.Operator_OPERATOR_UNSPECIFIED {
				return p.GetQuery(), fmt.Errorf("at ON: unknown operator")
			}

			p.pop()
			op2 := p.pop()
			op2split := strings.Split(op2, ".")
			if len(op2split) != 2 { //nolint:gomnd // ignore
				return p.GetQuery(), fmt.Errorf("at ON: expected <tablename>.<fieldname>")
			}

			currentCondition.RTable = op2split[0]
			currentCondition.ROperand = op2split[1]
			currentJoin := p.GetQuery().GetJoins()[len(p.GetQuery().GetJoins())-1]
			currentJoin.Conditions = append(currentJoin.GetConditions(), currentCondition)
			p.Query.Joins[len(p.GetQuery().GetJoins())-1] = currentJoin

			nextOp := p.peek()
			if strings.EqualFold(nextOp, "WHERE") {
				p.Step = Step_STEP_WHERE
			} else if strings.EqualFold(nextOp, "ORDER BY") {
				p.Step = Step_STEP_ORDER
			} else if strings.EqualFold(nextOp, "AND") {
				p.Step = Step_STEP_JOIN_CONDITION
			} else if strings.Contains(strings.ToUpper(nextOp), "JOIN") {
				p.Step = Step_STEP_JOIN
			}
		case Step_STEP_INSERT_FIELD_OPENING_PARENTS:
			openingParens := p.peek()
			if len(openingParens) != 1 || openingParens != "(" {
				return p.GetQuery(), fmt.Errorf("at INSERT INTO: expected opening parens")
			}

			p.pop()
			p.Step = Step_STEP_INSERT_FIELDS
		case Step_STEP_INSERT_FIELDS:
			identifier := p.peek()
			if !isIdentifier(identifier) {
				return p.GetQuery(), fmt.Errorf("at INSERT INTO: expected at least one field to insert")
			}

			p.Query.Fields = append(p.GetQuery().GetFields(), identifier)
			p.pop()
			p.Step = Step_STEP_INSERT_FIELDS_COMMA_OR_CLOSING_PARENTS
		case Step_STEP_INSERT_FIELDS_COMMA_OR_CLOSING_PARENTS:
			commaOrClosingParens := p.peek()
			if commaOrClosingParens != "," && commaOrClosingParens != ")" {
				return p.GetQuery(), fmt.Errorf("at INSERT INTO: expected comma or closing parens")
			}

			p.pop()
			if commaOrClosingParens == "," { //nolint:revive // false positive
				p.Step = Step_STEP_INSERT_FIELDS
				continue
			}

			p.Step = Step_STEP_INSERT_RWORD
		case Step_STEP_INSERT_RWORD:
			valuesRWord := p.peek()
			if !strings.EqualFold(valuesRWord, "VALUES") {
				return p.GetQuery(), fmt.Errorf("at INSERT INTO: expected 'VALUES'")
			}

			p.pop()
			p.Step = Step_STEP_INSERT_VALUES_OPENING_PARENS
		case Step_STEP_INSERT_VALUES_OPENING_PARENS:
			openingParens := p.peek()
			if openingParens != "(" {
				return p.GetQuery(), fmt.Errorf("at INSERT INTO: expected opening parens")
			}

			p.Query.Inserts = append(p.GetQuery().GetInserts(), &query.Query_Array{})
			p.pop()
			p.Step = Step_STEP_INSERT_VALUES
		case Step_STEP_INSERT_VALUES:
			quotedValue, ln := p.peekQuotedStringWithLength()
			if ln == 0 {
				return p.GetQuery(), fmt.Errorf("at INSERT INTO: expected quoted value")
			}

			p.Query.Inserts[len(p.GetQuery().GetInserts())-1].Items = append(p.GetQuery().GetInserts()[len(p.GetQuery().GetInserts())-1].GetItems(), quotedValue)
			p.pop()
			p.Step = Step_STEP_INSERT_VALUES_COMMA_OR_CLOSING_PARENS
		case Step_STEP_INSERT_VALUES_COMMA_OR_CLOSING_PARENS:
			commaOrClosingParens := p.peek()
			if commaOrClosingParens != "," && commaOrClosingParens != ")" {
				return p.GetQuery(), fmt.Errorf("at INSERT INTO: expected comma or closing parens")
			}

			p.pop()
			if commaOrClosingParens == "," {
				p.Step = Step_STEP_INSERT_VALUES
				continue
			}

			currentInsertRow := p.GetQuery().GetInserts()[len(p.GetQuery().GetInserts())-1]
			if len(currentInsertRow.GetItems()) < len(p.GetQuery().GetFields()) {
				return p.GetQuery(), fmt.Errorf("at INSERT INTO: value count doesn't match field count")
			}

			p.Step = Step_STEP_INSERT_VALUES_COMMA_BEFORE_OPENING_PARENS
		case Step_STEP_INSERT_VALUES_COMMA_BEFORE_OPENING_PARENS:
			commaRWord := p.peek()
			if commaRWord == ";" {
				p.Step = Step_STEP_SEMICOLON
			} else if !strings.EqualFold(commaRWord, ",") {
				return p.GetQuery(), fmt.Errorf("at INSERT INTO: expected comma")
			}

			p.pop()
			p.Step = Step_STEP_INSERT_VALUES_OPENING_PARENS
		case Step_STEP_CREATE_TABLE_NAME:
			identifier := p.peek()
			if !isIdentifier(identifier) {
				return p.GetQuery(), fmt.Errorf("at CREATE TABLE: table name cannot be empty")
			}

			p.pop()
			p.Query.TableName = identifier
			p.Step = Step_STEP_CREATE_TABLE_OPENING_PARENS
		case Step_STEP_CREATE_TABLE_OPENING_PARENS:
			openingParens := p.peek()
			p.pop()
			if openingParens != "(" {
				return p.GetQuery(), fmt.Errorf("at CREATE TABLE: expected opening parens")
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
				return p.GetQuery(), fmt.Errorf("at CREATE TABLE: expected at least one field to create table")
			}

			p.pop()
			typeField := p.peek()

			// get type field of table
			if !isIdentifier(typeField) {
				return p.GetQuery(), fmt.Errorf("at CREATE TABLE: expected at least one field to create table")
			}

			if p.GetQuery().GetTableFields() == nil {
				p.Query.TableFields = map[string]field.Type{}
			}

			// append field to table
			if vector.Contains(typeFieldTable, typeField) {
				switch typeField {
				case "int":
					p.Query.TableFields[identifier] = field.Type_TYPE_INTEGER
				case "integer":
					p.Query.TableFields[identifier] = field.Type_TYPE_INTEGER
				case "text":
					p.Query.TableFields[identifier] = field.Type_TYPE_STRING
				case "string":
					p.Query.TableFields[identifier] = field.Type_TYPE_STRING
				case "bool":
					p.Query.TableFields[identifier] = field.Type_TYPE_BOOLEAN
				case "boolean":
					p.Query.TableFields[identifier] = field.Type_TYPE_BOOLEAN
				default:
					return p.GetQuery(), fmt.Errorf("at CREATE TABLE: unsupported type of field")
				}

				p.pop()
			} else {
				return p.GetQuery(), fmt.Errorf("at CREATE TABLE: unsupported type of field")
			}

			p.Step = Step_STEP_CREATE_TABLE_FIELDS_COMMA_OR_CLOSING_PARENS
		case Step_STEP_CREATE_TABLE_FIELDS_COMMA_OR_CLOSING_PARENS:
			commaOrClosingParens := p.peek()
			if commaOrClosingParens != "," && commaOrClosingParens != ")" && commaOrClosingParens != ";" {
				return p.GetQuery(), fmt.Errorf("at CREATE TABLE: expected comma or closing parens")
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
		case Step_STEP_LIMIT:
			_ = p.peek()
			p.pop()

			countRaw := p.peek()
			p.pop()

			if countRaw == "" {
				return p.GetQuery(), fmt.Errorf("at LIMIT: empty LIMIT clause")
			}

			//nolint:revive // ignore this linter
			limit, err := strconv.ParseInt(countRaw, 10, 32)
			if err != nil {
				return p.GetQuery(), fmt.Errorf("at LIMIT: required number")
			}

			p.Query.Limit = int32(limit)
		case Step_STEP_CREATE_INDEX_NAME:
			if len(p.GetQuery().GetIndexs()) == 0 {
				p.Query.Indexs = []*v1.Index{}
			}

			// set name index of table
			p.Query.Indexs = append(p.GetQuery().GetIndexs(), &v1.Index{
				Name:   p.peek(),
				Type:   0,
				Fields: []string{},
			})
			p.pop()

			if p.peek() != "ON" {
				return p.GetQuery(), fmt.Errorf("at INDEX: incorrect sql-expression")
			}

			p.Step = Step_STEP_CREATE_INDEX_TABLE
		case Step_STEP_CREATE_INDEX_TABLE:
			_ = p.pop()
			// set name table
			p.Query.TableName = p.peek()
			p.pop()
			p.Step = Step_STEP_CREATE_INDEX_TYPE
		case Step_STEP_CREATE_INDEX_TYPE:
			_ = p.pop()
			// get type index
			switch strings.ToUpper(p.peek()) {
			case "BINARY":
				p.Query.Indexs[len(p.GetQuery().GetIndexs())-1].Type = v1.Type_TYPE_BINARY_SEARCH
			case "BTREE":
				p.Query.Indexs[len(p.GetQuery().GetIndexs())-1].Type = v1.Type_TYPE_BTREE
			case "HASH":
				p.Query.Indexs[len(p.GetQuery().GetIndexs())-1].Type = v1.Type_TYPE_HASH
			default:
				return p.GetQuery(), fmt.Errorf("at INDEX: incorrect type of index - %s", strings.ToUpper(p.peek()))
			}
			p.pop()

			p.Step = Step_STEP_CREATE_INDEX_PAYLOAD
		case Step_STEP_CREATE_INDEX_PAYLOAD:
			if p.peek() == "(" || p.peek() == "," {
				p.pop()
			}

			// set field for index
			p.Query.Indexs[len(p.GetQuery().GetIndexs())-1].Fields = append(p.GetQuery().GetIndexs()[len(p.GetQuery().GetIndexs())-1].GetFields(), p.peek())
			p.pop()

			if p.peek() == ")" {
				p.pop()
			}

			if p.peek() == ";" {
				p.Step = Step_STEP_SEMICOLON
				continue
			}

			p.Step = Step_STEP_CREATE_INDEX_PAYLOAD
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
	peeked, count := p.peekWithLength()
	p.I += count
	p.popWhitespace()

	return peeked
}

func (p *Parser) peekWithLength() (string, int32) {
	if p.GetI() >= int32(len(p.GetSql())) {
		return "", 0
	}

	for _, rWord := range reservedWords {
		token := strings.ToUpper(p.GetSql()[p.GetI():min(len(p.GetSql()), int(p.GetI())+len(rWord))])
		if token == rWord {
			return token, int32(len(token))
		}
	}

	if p.GetSql()[p.GetI()] == '\'' { // Quoted string
		return p.peekQuotedStringWithLength()
	}

	return p.peekIdentifierStringWithLength()
}

func (p *Parser) popWhitespace() {
	for ; p.GetI() < int32(len(p.GetSql())) && p.GetSql()[p.GetI()] == ' '; p.I++ {
	}
}

func (p *Parser) peekQuotedStringWithLength() (string, int32) {
	if int32(len(p.GetSql())) < p.GetI() || p.GetSql()[p.GetI()] != '\'' {
		return "", 0
	}

	for i := p.GetI() + 1; i < int32(len(p.GetSql())); i++ {
		if p.GetSql()[i] == '\'' && p.GetSql()[i-1] != '\\' {
			//nolint:gomnd // ignore
			return p.GetSql()[p.GetI()+1 : i], int32(len(p.GetSql()[p.GetI()+1:i]) + 2) // +2 for the two quotes
		}
	}

	return "", 0
}

func (p *Parser) peekIdentifierStringWithLength() (string, int32) {
	for i := p.GetI(); i < int32(len(p.GetSql())); i++ {
		if matched := r.MatchString(string(p.GetSql()[i])); !matched {
			return p.GetSql()[p.GetI():i], int32(len(p.GetSql()[p.GetI():i]))
		}
	}

	return p.GetSql()[p.GetI():], int32(len(p.GetSql()[p.GetI():]))
}

func (p *Parser) validate() error { //nolint:gocyclo,gocognit // ignore
	if p.GetQuery() == nil {
		return nil
	}

	if len(p.GetQuery().GetConditions()) == 0 && p.GetStep() == Step_STEP_WHERE_FIELD {
		return fmt.Errorf("at WHERE: empty WHERE clause")
	}

	if p.GetQuery().GetType() == query.Type_TYPE_UNSPECIFIED {
		return fmt.Errorf("query type cannot be empty")
	}

	if p.GetQuery().GetTableName() == "" {
		return fmt.Errorf("table name cannot be empty")
	}

	if len(p.GetQuery().GetConditions()) == 0 && (p.GetQuery().GetType() == query.Type_TYPE_UPDATE || p.GetQuery().GetType() == query.Type_TYPE_DELETE) {
		return fmt.Errorf("at WHERE: WHERE clause is mandatory for UPDATE & DELETE")
	}

	for _, c := range p.GetQuery().GetConditions() {
		if c.GetOperator() == query.Operator_OPERATOR_UNSPECIFIED {
			return fmt.Errorf("at WHERE: condition without operator")
		}

		if c.GetLValue() == "" && c.GetLValueIsField() {
			return fmt.Errorf("at WHERE: condition with empty left side operand")
		}

		if c.GetRValue() == "" && c.GetRValueIsField() {
			return fmt.Errorf("at WHERE: condition with empty right side operand")
		}
	}

	if p.GetQuery().GetType() == query.Type_TYPE_INSERT && len(p.GetQuery().GetInserts()) == 0 {
		return fmt.Errorf("at INSERT INTO: need at least one row to insert")
	}

	if p.GetQuery().GetType() == query.Type_TYPE_INSERT {
		for _, i := range p.GetQuery().GetInserts() {
			if len(i.GetItems()) != len(p.GetQuery().GetFields()) {
				return fmt.Errorf("at INSERT INTO: value count doesn't match field count")
			}
		}
	}

	return nil
}

func isIdentifier(s string) bool {
	for _, rw := range reservedWords {
		if strings.EqualFold(s, rw) {
			return false
		}
	}

	matched, _ := regexp.MatchString("[a-zA-Z_][a-zA-Z_0-9]*", s) //nolint:errcheck // ignore

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

func getOperator(operator string) query.Operator {
	switch operator {
	case "=":
		return query.Operator_OPERATOR_EQ
	case ">":
		return query.Operator_OPERATOR_GT
	case ">=":
		return query.Operator_OPERATOR_GTE
	case "<":
		return query.Operator_OPERATOR_LT
	case "<=":
		return query.Operator_OPERATOR_LTE
	case "!=":
		return query.Operator_OPERATOR_NE
	default:
		return query.Operator_OPERATOR_UNSPECIFIED
	}
}

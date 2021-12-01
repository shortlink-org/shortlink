package v1

import (
	"fmt"
	"regexp"
	"strings"

	v1 "github.com/batazor/shortlink/internal/pkg/database/query/v1"
)

var reservedWords = []string{
	"(", ")", ">=", ">", "<", "<=", "=", "!=", ",",
	"SELECT", "INSERT INTO", "VALUES", "UPDATE",
	"DELETE FROM", "WHERE", "FROM", "SET", "AS",
}

func New(sql string) (*Parser, error) {
	parser := &Parser{
		Sql: sql,
	}

	// Parse
	_, err := parser.Parse()
	if err != nil {
		return nil, err
	}

	return parser, nil
}

// Parse - main function that returns the "query struct" or an error
func (p *Parser) Parse() (*v1.Query, error) {
	// initial step
	//p.Step = "stepType"

	//for p.I < len(p.Sql) {
	//	nextToken := p.peek()
	//
	//	switch p.step {
	//	case stepType:
	//		switch nextToken {
	//		case UPDATE:
	//			p.Query.Type = "UPDATE"
	//			p.step = stepUpdateTable
	//		}
	//	case stepUpdateSet:
	//		continue
	//	case stepUpdateField:
	//		continue
	//	case stepUpdateComma:
	//		continue
	//	}
	//
	//	p.pop()
	//}

	return p.Query, nil
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
		if matched, _ := regexp.MatchString(`[a-zA-Z0-9]`, string(p.Sql[i])); !matched {
			return p.Sql[p.I:i], int32(len(p.Sql[p.I:i]))
		}
	}

	return p.Sql[p.I:], int32(len(p.Sql[p.I:]))
}

func (p *Parser) validate() error {
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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

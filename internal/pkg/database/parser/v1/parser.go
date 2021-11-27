package v1

import (
	v1 "github.com/batazor/shortlink/internal/pkg/database/query/v1"
)

var reserveWords = []string{
	"(", ")", ">=", ">", "<", "<=", "=", "!=", ",",
	"SELECT", "INSERT INTO", "VALUES", "UPDATE",
	"DELETE FROM", "WHERE", "FROM", "SET",
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
	peeked, _ := p.peekWithLength()
	//p.I += len
	p.popWhitespace()
	return peeked
}

func (p *Parser) peekWithLength() (string, int) {
	if p.I >= int32(len(p.Sql)) {
		return "", 0
	}

	return p.peekIdentifierStringWithLength()
}

func (p *Parser) popWhitespace() {

}

func (p *Parser) peekQuotedStringWithLength() (string, int) {
	return "", 0
}

func (p *Parser) peekIdentifierStringWithLength() (string, int) {
	return "", 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

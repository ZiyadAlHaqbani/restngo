package parser

import (
	"htestp/dsl/scanner"
	"htestp/nodes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {

	success_tokens := []scanner.Token{
		{Type: scanner.Node, Content: "StaticNode"},
		{Type: scanner.LeftParen, Content: "("},
		{Type: scanner.StringLiteral, Content: "test"},
		{Type: scanner.Comma, Content: ","},
		{Type: scanner.METHOD, Content: "GET"},
		{Type: scanner.Comma, Content: ","},
		{Type: scanner.StringLiteral, Content: "/test"},
		{Type: scanner.Comma, Content: ","},
		{Type: scanner.Constraint, Content: "ExistConstraint"},
		{Type: scanner.LeftParen, Content: "("},
		{Type: scanner.StringLiteral, Content: "test"},
		{Type: scanner.Comma, Content: ","},
		{Type: scanner.Float64, Content: "FLOAT64"},
		{Type: scanner.RightParen, Content: ")"},
		{Type: scanner.RightParen, Content: ")"},
	}

	successful := CreateParser(success_tokens)
	successful.Parse()

	var expected *nodes.StaticNode
	assert.IsType(t, expected, successful.Head, "expected type StaticNode")

	failed_tokens := []scanner.Token{
		{Type: scanner.Node, Content: "TestNode"},
		{Type: scanner.LeftParen, Content: "("},
		{Type: scanner.StringLiteral, Content: "test.field[12]"},
		{Type: scanner.Comma, Content: ","},
		{Type: scanner.METHOD, Content: "GET"},
		{Type: scanner.Comma, Content: ","},
		{Type: scanner.StringLiteral, Content: "/test"},
		{Type: scanner.Comma, Content: ","},
		{Type: scanner.Constraint, Content: "ExistConstraint"},
		{Type: scanner.LeftParen, Content: "("},
		{Type: scanner.StringLiteral, Content: "test"},
		{Type: scanner.Comma, Content: ","},
		{Type: scanner.METHOD, Content: "FLOAT64"},
		{Type: scanner.Comma, Content: ","},
		{Type: scanner.StringLiteral, Content: "10.5435"},
		{Type: scanner.Comma, Content: ","},
		{Type: scanner.RightParen, Content: ")"},
		{Type: scanner.RightParen, Content: ")"},
	}

	failed := CreateParser(failed_tokens)
	assert.Panics(t, func() {
		failed.Parse()
	}, "expected panic for incorrect token sequence")

}

func TestParseConstraint(t *testing.T) {

	success_tokens := []scanner.Token{
		{Type: scanner.Constraint, Content: "ExistConstraint"},
		{Type: scanner.LeftParen, Content: "("},
		{Type: scanner.StringLiteral, Content: "test.field[12]"},
		{Type: scanner.Comma, Content: ","},
		{Type: scanner.Float64, Content: "FLOAT64"},
		{Type: scanner.RightParen, Content: ")"},
	}
	_ = success_tokens
	successful := CreateParser(success_tokens)
	successful.parseConstraint()

}

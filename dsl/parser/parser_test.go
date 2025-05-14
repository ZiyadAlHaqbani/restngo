package parser

import (
	"htestp/constraints"
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

func TestParseExistConstraint(t *testing.T) {

	success_tokens := []scanner.Token{
		{Type: scanner.Constraint, Content: "ExistConstraint"},
		{Type: scanner.LeftParen, Content: "("},
		{Type: scanner.StringLiteral, Content: "test.field[12]"},
		{Type: scanner.Comma, Content: ","},
		{Type: scanner.Float64, Content: "FLOAT64"},
		{Type: scanner.Comma, Content: ","},
		{Type: scanner.RightParen, Content: ")"},
	}
	_ = success_tokens
	successful := CreateParser(success_tokens)
	var expected *constraints.Exist_Constraint
	assert.IsType(t, expected, successful.parseConstraint(), "expected constraints.Exist_Constraint")

	fail_tokens := []scanner.Token{
		{Type: scanner.Constraint, Content: "ExistStoreConstraint"},
		{Type: scanner.LeftParen, Content: "("},
		{Type: scanner.StringLiteral, Content: "test.field[12]"},
		{Type: scanner.Comma, Content: ","},
		{Type: scanner.Float64, Content: "FLOAT64"},
		// {Type: scanner.Comma, Content: ","},	// no comma between args
		{Type: scanner.StringLiteral, Content: "varname"},
		{Type: scanner.Comma, Content: ","},
		{Type: scanner.RightParen, Content: ")"},
	}

	failed := CreateParser(fail_tokens)
	assert.Panics(t, func() { failed.parseConstraint() }, "expected parser to panic")

}

func TestParseExistStoreConstraint(t *testing.T) {

	success_tokens := []scanner.Token{
		{Type: scanner.Constraint, Content: "ExistStoreConstraint"},
		{Type: scanner.LeftParen, Content: "("},
		{Type: scanner.StringLiteral, Content: "test.field[12]"},
		{Type: scanner.Comma, Content: ","},
		{Type: scanner.Float64, Content: "FLOAT64"},
		{Type: scanner.Comma, Content: ","},
		{Type: scanner.StringLiteral, Content: "test_varname"},
		{Type: scanner.RightParen, Content: ")"},
	}
	_ = success_tokens
	successful := CreateParser(success_tokens)
	successful_constraint := successful.parseConstraint()
	var expected *constraints.Exist_Store_Constraint
	assert.IsType(t, expected, successful_constraint, "expected Exist_Store_Constraint")

	fail_tokens := []scanner.Token{
		{Type: scanner.Constraint, Content: "ExistStoreConstraint"},
		{Type: scanner.LeftParen, Content: "("},
		{Type: scanner.StringLiteral, Content: "test.field[12]"},
		{Type: scanner.Comma, Content: ","},
		{Type: scanner.Float64, Content: "FLOAT64"},
		// {Type: scanner.Comma, Content: ","}, removed a comma after the type
		{Type: scanner.StringLiteral, Content: "test_varname"},
		{Type: scanner.RightParen, Content: ")"},
	}

	failed := CreateParser(fail_tokens)
	assert.Panics(t, func() { failed.parseConstraint() }, "expected parser to panic")

}

func TestParseMatchConstraint(t *testing.T) {
	//TODO: implement match in the parser first
	t.SkipNow()
}

func TestParseMatchStoreConstraint(t *testing.T) {
	//TODO: implement match in the parser first
	t.SkipNow()
}

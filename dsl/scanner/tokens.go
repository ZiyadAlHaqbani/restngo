package scanner

import (
	"htestp/models"
	"strconv"
)

type TokenType int

const (
	EOF TokenType = iota
	Comma
	LeftParen
	RightParen

	//	Value wrappers
	Number
	StringLiteral

	//	Data types
	Float64
	BOOL
	STRING
	ARRAY
	OBJECT

	Identifier //TODO: remove the below types, as identifier is handled by the parser

	Node
	Constraint

	URL
	METHOD
)

var TokenTypeToString = map[TokenType]string{
	EOF:           "EOF",
	Comma:         "Comma",
	LeftParen:     "LeftParen",
	RightParen:    "RightParen",
	Number:        "Number",
	StringLiteral: "StringLiteral",
	Float64:       "Float64",
	BOOL:          "BOOL",
	STRING:        "STRING",
	ARRAY:         "ARRAY",
	OBJECT:        "OBJECT",
	Identifier:    "Identifier",
	Node:          "Node",
	Constraint:    "Constraint",
	URL:           "URL",
	METHOD:        "METHOD",
}

// Map strings to specific token types
var TypesMap = map[string]TokenType{
	"StaticNode":      Node,
	"DynamicNode":     Node,
	"ConditionalNode": Node,

	"ExistConstraint":      Constraint,
	"ExistStoreConstraint": Constraint,
	"MatchConstraint":      Constraint,
	"MatchStoreConstraint": Constraint,

	"Float64": Float64,
	"BOOL":    BOOL,
	"STRING":  STRING,
	"ARRAY":   ARRAY,
	"OBJECT":  OBJECT,

	"GET":  METHOD,
	"POST": METHOD,
	"PUT":  METHOD,
}

// Map TokenTypes to specific type enums
var DataTypesMap = map[TokenType]models.MatchType{

	Float64: models.TypeFloat,
	BOOL:    models.TypeBool,
	STRING:  models.TypeString,
	ARRAY:   models.TypeArray,
	OBJECT:  models.TypeObject,
}

type Token struct {
	Content string
	Type    TokenType
	Start   int
	Line    int
}

func (token Token) ToString() string {
	return "content: " + token.Content + ", type: " + TokenTypeToString[token.Type] +
		", line: " + strconv.Itoa(token.Line) + ", start at: " + strconv.Itoa(token.Start)
}

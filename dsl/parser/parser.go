package parser

import (
	"htestp/constraints"
	"htestp/dsl/scanner"
	httphandler "htestp/http_handler"
	"htestp/models"
	"htestp/nodes"
	"log"
	"slices"
	"strconv"
)

func CreateParser(tokens []scanner.Token) *Parser {
	return &Parser{
		tokens: tokens,
	}
}

type Parser struct {
	tokens  []scanner.Token
	current int

	CURRENT_TOKEN scanner.Token
	Head          models.Node
}

func (parser *Parser) peek() scanner.Token {
	parser.CURRENT_TOKEN = parser.tokens[parser.current]
	return parser.tokens[parser.current]
}

func (parser *Parser) advance() scanner.Token {
	temp := parser.peek()
	parser.current++
	return temp
}

func (parser *Parser) check(t scanner.TokenType) bool {
	return parser.peek().Type == t
}

func (parser *Parser) match(t ...scanner.TokenType) bool {
	if slices.Contains(t, parser.peek().Type) {
		parser.advance()
		return true
	}
	return false

}

func (parser *Parser) consume(t scanner.TokenType) scanner.Token {
	if parser.peek().Type == t {
		return parser.advance()
	}

	log.Panicf("ERROR: expected %s, but found: '%s'", scanner.TokenTypeToString[t], parser.peek().ToString())
	return scanner.Token{}
}

func (parser *Parser) Parse() {
	parser.Head = parser.parseExpression()
}

func (parser *Parser) parseExpression() models.Node {
	return parser.parseFunction()
}

func (parser *Parser) parseFunction() models.Node {
	if scanner.TypesMap[parser.peek().Content] == scanner.Node {
		return parser.parseNode()
	} else if scanner.TypesMap[parser.peek().Content] == scanner.Constraint {
		log.Panic("ERROR: constraints must only be defined within nodes, and not outside of them")
	}

	log.Panicf("ERROR: unrecognized identifier: %+v", parser.peek())
	//_
	return &nodes.ConditionalNode{}
}

func (parser *Parser) parseNode() models.Node {

	id := parser.advance()
	switch id.Content {
	case "StaticNode":
		return parser.parseStaticNode()
	case "DynamicNode":
		log.Panic("ERROR: DynamicNode is not supported yet.")
	default:
		log.Panicf("ERROR: unsupported node type: %q, at: %+v", id.Content, id)

	}

	log.Panic("ERROR: check code!")
	//_
	return &nodes.ConditionalNode{}
}

func (parser *Parser) parseStaticNode() models.Node {
	//	identifier already consumed
	parser.consume(scanner.LeftParen)

	staticNode := nodes.StaticNode{}

	//TODO: find a way to support node IDs through the DSL
	id := parser.consume(scanner.StringLiteral)
	parser.consume(scanner.Comma)
	_ = id

	method := parser.consume(scanner.METHOD)
	parser.consume(scanner.Comma)

	url := parser.consume(scanner.StringLiteral)
	parser.consume(scanner.Comma)

	staticNode.Request = httphandler.Request{
		Url:    url.Content,
		Method: method.Content,
	}

	constraints := parser.parseConstraints()

	staticNode.Constraints = constraints

	children := []models.Node{}
	for parser.check(scanner.Node) {
		new := parser.parseNode()
		children = append(children, new)
		commaFound := parser.match(scanner.Comma)
		_ = commaFound
	}
	staticNode.Next = children

	parser.consume(scanner.RightParen)

	//_
	return &staticNode
}

func (parser *Parser) parseConstraints() []models.Constraint {
	var constraints []models.Constraint

	for parser.check(scanner.Constraint) {

		new := parser.parseConstraint()
		if new == nil {
			break
		}
		constraints = append(constraints, new)

	}

	return constraints
}

func (parser *Parser) parseConstraint() models.Constraint {

	var toReturn models.Constraint

	identifier := parser.advance()
	parser.match(scanner.LeftParen)

	field := parser.consume(scanner.StringLiteral)
	parser.consume(scanner.Comma)

	expected, exists := scanner.DataTypesMap[parser.advance().Type]
	if !exists {
		log.Panicf("ERROR: unrecognized type: %v", parser.advance().Type)
	}

	switch identifier.Content {
	case "ExistConstraint":
		toReturn = &constraints.Exist_Constraint{
			Field: field.Content,
			Type:  expected,
		}
	case "ExistStoreConstraint":
		varname := parser.consume(scanner.StringLiteral)
		toReturn = &constraints.Exist_Store_Constraint{
			InnerConstraint: constraints.Exist_Constraint{},
			Varname:         varname.Content,
		}
	case "MatchConstraint":
		temp := constraints.Match_Constraint{
			Field:    field.Content,
			Type:     expected,
			Expected: nil,
			Status:   models.MatchStatus{},
		}

		switch expected {
		case models.TypeFloat:
			num := parser.consume(scanner.Number)
			val, err := strconv.ParseFloat(num.Content, 64)
			if err != nil {
				log.Panicf("ERROR: couldn't convert %q into a number: %+v", num.Content, err)
			}
			temp.Expected = val
		case models.TypeString:
			expected_val := parser.consume(scanner.StringLiteral)
			temp.Expected = expected_val.Content
		}
		toReturn = &temp
	case "MatchStoreConstraint":
		varname := parser.consume(scanner.StringLiteral)
		toReturn = &constraints.Exist_Store_Constraint{
			InnerConstraint: constraints.Exist_Constraint{},
			Varname:         varname.Content,
		}
	}

	parser.consume(scanner.RightParen)
	parser.match(scanner.Comma)

	return toReturn

}

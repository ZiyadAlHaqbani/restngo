package parser

import (
	"htestp/constraints"
	"htestp/dsl/scanner"
	httphandler "htestp/http_handler"
	"htestp/models"
	"htestp/nodes"
	"log"
	"slices"
)

func CreateParser(tokens []scanner.Token) *Parser {
	return &Parser{
		tokens: tokens,
	}
}

type Parser struct {
	tokens  []scanner.Token
	current int

	head models.Node
}

func (parser *Parser) peek() scanner.Token {
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
	log.Fatalf("ERROR: expected %s, but found: '%s'", scanner.TokenTypeToString[t], parser.peek().ToString())
	return scanner.Token{}
}

func (parser *Parser) Parse() {
	parser.head = parser.parseExpression()
}

func (parser *Parser) parseExpression() models.Node {
	return parser.parseFunction()
	//_
	return &nodes.ConditionalNode{}
}

func (parser *Parser) parseFunction() models.Node {
	if scanner.TypesMap[parser.peek().Content] == scanner.Node {
		return parser.parseNode()
	} else if scanner.TypesMap[parser.peek().Content] == scanner.Constraint {
		log.Fatal("ERROR: constraints must only be defined within nodes, and not outside of them")
	}

	log.Fatalf("ERROR: unrecognized identifier: %+v", parser.peek())
	//_
	return &nodes.ConditionalNode{}
}

func (parser *Parser) parseNode() models.Node {

	id := parser.advance()
	switch id.Content {
	case "StaticNode":
		return parser.parseStaticNode()
	case "DynamicNode":
		log.Fatal("ERROR: DynamicNode is not supported yet.")
	default:
		log.Fatalf("ERROR: unsupported node type: %q, at: %+v", id.Content, id)

	}

	log.Fatal("ERROR: check code!")
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

	parser.parseNode() //FDFDFSDFSD

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

	identifier := parser.advance()
	parser.match(scanner.LeftParen)

	field := parser.consume(scanner.StringLiteral)
	parser.consume(scanner.Comma)

	expected, exists := scanner.DataTypesMap[parser.advance().Type]
	if !exists {
		log.Fatalf("ERROR: unrecognized type: %v", parser.advance().Type)
	}

	switch identifier.Content {
	case "ExistConstraint":
		return &constraints.Exist_Constraint{
			Field: field.Content,
			Type:  expected,
		}
	}
	if !parser.match(scanner.Comma) {
		return nil
	}

	log.Fatalf("ERROR: couldn't parse constraint for token sequence starting at: %+v", identifier)
	return nil
}

package scanner

type TokenType int

const (
	INVALID TokenType = iota
	Comma
	LeftParen
	RightParen

	Number
	StringLiteral
	DataType

	StaticNode
	//TODO: support other node types

	ExistConstraint
	//TODO: support other constraint types

	URL
	METHOD
)

// Map strings to specific token types
var TypesMap = map[string]TokenType{
	"StaticNode":      StaticNode,
	"ExistConstraint": ExistConstraint,
	"GET":             METHOD,
	"POST":            METHOD,
	"PUT":             METHOD,
}

type Token struct {
	Content string
	Type    TokenType

	Line int
}

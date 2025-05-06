package scanner

import (
	"log"
	"slices"
)

func CreateScanner(source string) Scanner {
	return Scanner{
		start:   0,
		current: 0,
		line:    0,
		source:  source,
		tokens:  []Token{},
	}
}

type Scanner struct {
	start   int
	current int
	line    int

	source string //TODO: use a type that only stores characters
	tokens []Token
}

func (scanner *Scanner) Scan() []Token {

	for !scanner.eof() {
		scanner.start = scanner.current

		scanner.scanToken()

	}
	scanner.addToken(EOF)

	return scanner.tokens
}

func (scanner *Scanner) scanToken() {
	switch scanner.peek() {

	//	skip white space
	case ' ', '\r', '\t':
		scanner.advance()

	// new lines
	case '\n':
		scanner.advance()
		scanner.line++

	case ',':
		scanner.advance()
		scanner.addToken(Comma)
	case '(':
		scanner.advance()
		scanner.addToken(LeftParen)
	case ')':
		scanner.advance()
		scanner.addToken(RightParen)

	case '"':
		scanner.stringLiteral()

	default:
		if isDigit(scanner.peek()) {
			scanner.number()
		} else if isAlpha(scanner.peek()) {
			scanner.identifier()
		}
	}

}

func (scanner *Scanner) peek() byte {
	if scanner.eof() {
		//'\0'
		return 0
	} else {
		return scanner.source[scanner.current]
	}
}

func (scanner *Scanner) peekAhead() byte {
	if scanner.current+1 >= len(scanner.source) {
		//'\0'
		return 0
	} else {
		return scanner.source[scanner.current+1]
	}
}

func (scanner *Scanner) advance() byte {
	temp := scanner.source[scanner.current]
	scanner.current++
	return temp
}

//	tries to match the current character with any character in the given list, returns
//
// true if successful
func (scanner *Scanner) match(chars ...byte) bool {

	if slices.Contains(chars, scanner.peek()) {
		scanner.current++
		return true
	}

	return false
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isAlpha(c byte) bool {
	return c >= 'a' && c <= 'z' ||
		c >= 'A' && c <= 'Z' ||
		c == '_'
}

func isAlphaNumeric(c byte) bool {
	return isAlpha(c) || isDigit(c)
}

func (scanner *Scanner) stringLiteral() {
	currChar := string(scanner.peek())
	_ = currChar
	scanner.advance()
	for !scanner.eof() && !scanner.match('"') {
		scanner.advance()
	}

	tok := scanner.addToken(StringLiteral)

	if scanner.eof() {
		log.Panicf("ERROR: string literals must end with a '\"', at: %+v", tok)
	}

}

func (scanner *Scanner) identifier() {

	for !scanner.eof() && isAlphaNumeric(scanner.peek()) {
		scanner.advance()
	}

	tok := Token{
		Content: scanner.source[scanner.start:scanner.current],
		Type:    0,
		Line:    scanner.line,
	}

	if scanner.eof() {
		log.Panicf("ERROR: string literals must end with a '\"', at: %+v", tok)
	}

	tokType, exists := TypesMap[tok.Content]
	if !exists {
		log.Panicf("ERROR: unrecognized literal: %q at: %+v", tok.Content, tok)
	}

	tok.Type = tokType

	scanner.addTokenRaw(tok)
}

func (scanner *Scanner) number() {
	for isDigit(scanner.peek()) {
		scanner.advance()
	}

	if !isWhiteSpace(scanner.peek()) && scanner.peek() != ',' {
		log.Panicf("ERROR: numbers must only end with whitespace or comma ',', at: %+v", scanner.addToken(Number))
	}

	scanner.addToken(Number)
}

func (scanner *Scanner) eof() bool {
	return scanner.current >= len(scanner.source)
}

func (scanner *Scanner) addToken(Type TokenType) Token {
	token := Token{
		Type:  Type,
		Start: scanner.current,
		Line:  scanner.line,
	}

	if Type == StringLiteral {
		temp := scanner.source[scanner.start:scanner.current]
		temp = temp[1:]
		temp = temp[:len(temp)-1]
		token.Content = temp
	} else {
		token.Content = scanner.source[scanner.start:scanner.current]
	}

	scanner.tokens = append(scanner.tokens, token)

	return token
}

func (scanner *Scanner) addTokenRaw(tok Token) Token {

	scanner.tokens = append(scanner.tokens, tok)

	return tok
}

func (scanner *Scanner) ToString() string {
	temp := ""
	for _, tok := range scanner.tokens {
		temp += tok.Content
	}
	return temp
}

func isWhiteSpace(c byte) bool {

	return c == ' ' || c == '\n' || c == '\r' || c == '\t'

}

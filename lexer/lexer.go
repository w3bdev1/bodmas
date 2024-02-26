package lexer

import "github.com/w3bdev1/bodmas/token"

type Lexer struct {
	Content      string
	NextPosition int
	Char         byte
}

func NewLexer(content string) *Lexer {
	l := &Lexer{
		Content:      content,
		NextPosition: 0,
	}

	l.ReadChar()
	return l
}

func (l *Lexer) ReadChar() {
	if l.NextPosition >= len(l.Content) {
		l.Char = 0
	} else {
		l.Char = l.Content[l.NextPosition]
	}
	l.NextPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	l.skipWhitespace()
	var tok token.Token
	switch l.Char {
	case '(':
		tok = createToken(token.LPAREN, l.Char)
	case ')':
		tok = createToken(token.RPAREN, l.Char)
	case '+':
		tok = createToken(token.ADD, l.Char)
	case '-':
		tok = createToken(token.SUBTRACT, l.Char)
	case '*':
		tok = createToken(token.MULTIPLY, l.Char)
	case '/':
		tok = createToken(token.DIVIDE, l.Char)
	case '^':
		tok = createToken(token.EXPONENT, l.Char)
	case 0:
		tok.Type = token.EOF
		tok.Value = ""
	default:
		if isDigit(l.Char) {
			tok.Type = token.INT
			tok.Value = l.getNum()
			return tok
		} else {
			tok = createToken(token.ILLEGAL, l.Char)
		}
	}

	l.ReadChar()
	return tok
}

func createToken(tokenType string, value byte) token.Token {
	return token.Token{Type: tokenType, Value: string(value)}
}

func isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}

func (l *Lexer) getNum() string {

	start_location := l.NextPosition - 1
	for isDigit(l.Char) {
		l.ReadChar()
	}
	end_location := l.NextPosition - 1

	return l.Content[start_location:end_location]
}

func (l *Lexer) skipWhitespace() {
	for l.Char == ' ' {
		l.ReadChar()
	}
}

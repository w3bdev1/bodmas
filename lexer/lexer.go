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
		l.NextPosition += 1
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	switch l.Char {
	case '+':
		tok = createToken(token.ADD, l.Char)
	case '-':
		tok = createToken(token.SUBTRACT, l.Char)
	case '*':
		tok = createToken(token.MULTIPLY, l.Char)
	case '/':
		tok = createToken(token.DIVIDE, l.Char)
	default:
		if isDigit(l.Char) {
			tok.Type = token.INT
			tok.Value = l.getNum()
			return tok
		}
	}

	l.ReadChar()
	return tok
}

func createToken(tokenType string, value byte) token.Token {
	return token.Token{Type: tokenType, Value: string(value)}
}

func isDigit(char byte) bool {
	return char >= 0 && char <= 9
}

func (l *Lexer) getNum() string {

	start_location := l.NextPosition - 1
	for isDigit(l.Char) {
		l.ReadChar()
	}
	end_location := l.NextPosition - 1

	return l.Content[start_location:end_location]
}

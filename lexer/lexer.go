package lexer

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

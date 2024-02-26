package token

type Token struct {
	Type  string
	Value string
}

/* token types */
const (
	LPAREN = "("
	RPAREN = ")"

	ADD      = "+"
	SUBTRACT = "-"
	MULTIPLY = "*"
	DIVIDE   = "/"
	EXPONENT = "^"

	INT = "INT"
	// TODO: FLOAT = "FLOAT"

	EOF     = "EOF"
	ILLEGAL = "ILLEGAL"
)

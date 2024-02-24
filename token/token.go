package token

type Token struct {
	Type  string
	Value string
}

/* token types */
const (
	ADD      = "+"
	SUBTRACT = "-"
	MULTIPLY = "*"
	DIVIDE   = "/"
	// TODO: EXPONENT = "^"

	INT = "INT"
	// TODO: FLOAT = "FLOAT"
)

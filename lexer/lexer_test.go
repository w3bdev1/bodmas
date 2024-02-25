package lexer_test

import (
	"fmt"
	"testing"

	"github.com/w3bdev1/bodmas/lexer"
	"github.com/w3bdev1/bodmas/token"
)

func TestNewLexer(t *testing.T) {
	content := "abcd"

	l := lexer.NewLexer(content)

	if !(l.Content == content && l.NextPosition == 1 && l.Char == content[0]) {
		t.Fatalf("NewLexer(%s) creation failed:", content)
	}
}

func TestNextToken(t *testing.T) {
	content := "20 + 4 * 3 / 6 - 1"

	expectations := []struct {
		expectedType  string
		expectedValue string
	}{
		{token.INT, "20"},
		{token.ADD, "+"},
		{token.INT, "4"},
		{token.MULTIPLY, "*"},
		{token.INT, "3"},
		{token.DIVIDE, "/"},
		{token.INT, "6"},
		{token.SUBTRACT, "-"},
		{token.INT, "1"},
	}

	l := lexer.NewLexer(content)

	for index, test := range expectations {
		tok := l.NextToken()
		fmt.Println(test)
		fmt.Println(tok)
		fmt.Println(l)

		if tok.Type != test.expectedType {
			t.Fatalf("Test[%d] failed: expected type %s, got %s", index, test.expectedType, tok.Type)
		}

		if tok.Value != test.expectedValue {
			t.Fatalf("Test[%d] failed: expected value %s, got %s", index, test.expectedValue, tok.Value)
		}
	}
}

package lexer_test

import (
	"testing"

	"github.com/w3bdev1/bodmas/lexer"
)

func TestNewLexer(t *testing.T)  {
  content := "abcd"

  l := lexer.NewLexer(content)

  if !(l.Content == content && l.NextPosition == 1 && l.Char == content[0]) {
     t.Fatalf("NewLexer(%s) creation failed:", content)
  }
}

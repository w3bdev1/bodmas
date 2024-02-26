package main

import (
	"os"

	"github.com/w3bdev1/bodmas/repl"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}

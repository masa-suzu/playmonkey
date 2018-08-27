package main

import (
	"os"

	"github.com/masa-suzu/monkey/repl"
)

func main() {
	repl.Start(os.Stdin, os.Stdout, ">> ")
}

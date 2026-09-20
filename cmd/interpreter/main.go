package main

import (
	"fmt"
	"os"

	"arkana/pkg/scanner"
)

func fail(format string, arguments ...any) {
	fmt.Fprintf(os.Stderr, "lab1: "+format+"\n", arguments...)
	os.Exit(65)
}

func main() {
	if len(os.Args) < 2 {
		fail("expected '--tokenize <source-file>'")
	}

	if os.Args[1] == "--tokenize" {
		if len(os.Args) < 3 {
			fail("expected source-file path after '--tokenize'")
		}

		path := os.Args[2]
		source, err := os.ReadFile(path)
		if err != nil {
			fail("cannot read '%s': %v", path, err)
		}

		s := scanner.NewScanner(string(source))
		tokens := s.ScanTokens()

		if s.IsError {
			os.Exit(65)
		}

		for _, token := range tokens {
			fmt.Printf("Token(type=%v, lexeme=%s, literal=%v, line=%d)\n",
				token.Type,
				token.Lexeme,
				scanner.StringLiteral(token.Literal),
				token.Line,
			)
		}

		return
	}

	// Lab 0 behavior
	path := os.Args[1]
	source, err := os.ReadFile(path)
	if err != nil {
		fail("cannot read '%s': %v", path, err)
	}

	if _, err := os.Stdout.Write(source); err != nil {
		fmt.Fprintf(os.Stderr, "lab0: cannot write output: %v\n", err)
		os.Exit(70)
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"

	"arkana/pkg/parser"
	"arkana/pkg/scanner"
)

func fail(format string, arguments ...any) {
	fmt.Fprintf(os.Stderr, "lab1: "+format+"\n", arguments...)
	os.Exit(65)
}

func printTokens(tokens []scanner.Token) {
	for _, token := range tokens {
		fmt.Printf("Token(type=%v, lexeme=%s, literal=%v, line=%d)\n",
			token.Type,
			token.Lexeme,
			scanner.StringLiteral(token.Literal),
			token.Line,
		)
	}
}

func runREPL() {
	reader := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("> ")

		if !reader.Scan() {
			break
		}

		source := reader.Text()

		s := scanner.NewScanner(source)
		tokens := s.ScanTokens()
		printTokens(tokens)
	}

	// check for input errors after the REPL loop ends
	if err := reader.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "lab1: error reading input: %v\n", err)
	}
}

func main() {
	if len(os.Args) == 1 {
		runREPL()
		return
	}

	// Lab 1
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

		printTokens(tokens)

		return
	}

	// Lab 2
	if os.Args[1] == "--parse" {
		if len(os.Args) < 3 {
			fail("expected source-file path after '--parse'")
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

		p := parser.NewParser(tokens)
		expression := p.Parse()

		// AST printer will go here
		//_ = expression

		// temporary debug until printer is implemented
		fmt.Printf("Type: %T, Value: %#v\n", expression, expression)

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

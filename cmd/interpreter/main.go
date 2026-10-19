package main

import (
	"bufio"
	"fmt"
	"os"

	"arkana/pkg/ast"
	"arkana/pkg/interpreter"
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
		// printTokens(tokens) previously for scanner lab

		if s.IsError {
			continue // scanner already printed its error
		}

		if len(tokens) == 1 {
			continue // just EOF = blank line or only a comment
		}

		p := parser.NewParser(tokens)
		tree := p.Parse() // p.Parse() returns a Tree whose children are ASTs

		if p.IsError {
			continue // parser already printed its error
		}

		for _, node := range tree.ASTs {
			fmt.Println(ast.Print(node))
		}
	}

	// check for input errors after the REPL loop ends
	if err := reader.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading input: %v\n", err)
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
		tree := p.Parse() // p.Parse() returns a Tree whose children are the program's ASTs

		if p.IsError {
			os.Exit(65)
		}

		for _, node := range tree.ASTs {
			fmt.Println(ast.Print(node))
		}

		return
	}

	// Lab 3
	if os.Args[1] == "--eval" {
		if len(os.Args) < 3 {
			fail("expected source-file path after '--eval'")
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
		tree := p.Parse()

		if p.IsError {
			os.Exit(65)
		}

		for _, node := range tree.ASTs {
			fmt.Println(interpreter.Stringify(interpreter.Evaluate(node)))
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

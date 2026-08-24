package main

import (
	"fmt"
	"os"
)

func fail(format string, arguments ...any) {
	fmt.Fprintf(os.Stderr, "lab0: "+format+"\n", arguments...)
	os.Exit(65)
}

func main() {
	if len(os.Args) < 2 {
		fail("expected one source-file path")
	}

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
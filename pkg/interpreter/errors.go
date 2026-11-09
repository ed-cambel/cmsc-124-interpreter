package interpreter

import (
	"fmt"

	"arkana/pkg/scanner"
)

// RuntimeError is the only error the evaluator throws on purpose
type RuntimeError struct {
	Token   scanner.Token
	Message string
}

func (e RuntimeError) Error() string {
	return fmt.Sprintf("[Line %d] Runtime error: %s", e.Token.Line, e.Message)
}

// evaluates one expression and is the only place that recovers
func catch(err *error) {
	if r := recover(); r != nil {
		re, ok := r.(RuntimeError) // runtime error type
		if !ok {
			panic(r) // e.g. the "Unknown statement type." string, re-panic
		}
		*err = re
	}
}

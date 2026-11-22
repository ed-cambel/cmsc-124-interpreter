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

// recovers a panic raised during evaluation and stores it in *err
// deferred by Execute and Interpret, it is the only place that recovers
func catch(err *error) {
	r := recover()
	if r == nil {
		return
	}

	re, isRe := r.(RuntimeError)
	if !isRe {
		// a bug in the interpreter itself
		*err = fmt.Errorf("Runtime error: internal error: %v", r)
		return
	}
	*err = re
}

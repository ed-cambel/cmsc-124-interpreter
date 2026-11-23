package interpreter

import (
	"arkana/pkg/scanner"
)

type kind int

const (
	kindNil kind = iota
	kindGrain
	kindDram
	kindScroll
	kindOther
	kindSpell
)

// maps a declaration's type keyword to the kind its variable is checked against
func declaredKind(t scanner.TokenType) kind {
	switch t {
	case scanner.GRAIN:
		return kindGrain
	case scanner.DRAM:
		return kindDram
	case scanner.SCROLL:
		return kindScroll
	}
	return kindNil // seal: Define ignores this and infers the type from the value
}

func kindOf(v any) kind {
	switch v.(type) {
	case nil:
		return kindNil
	case int:
		return kindGrain
	case float64:
		return kindDram
	case string:
		return kindScroll	
	case Callable:
		return kindSpell
	}
	return kindOther
}


func coerceKind(declared kind, v any) (any, bool) {
	kind := kindOf(v)

	switch {
	case kind == kindNil, kind == declared:
		return v, true
	case declared == kindDram && kind == kindGrain:
		return float64(v.(int)), true
	}

	return nil, false
}
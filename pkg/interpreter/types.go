package interpreter

type kind int

const (
	kindNil kind = iota
	kindGrain
	kindDram
	kindScroll
	kindOther
)

func kindOf(v any) kind {
	switch v.(type) {
	case nil:
		return kindNil
	case int64:
		return kindGrain
	case float64:
		return kindDram
	case string:
		return kindScroll
	default:
		return kindOther
	}
}

func coerceKind(v any, declared kind) (any, bool) {
	kind := kindOf(v)

	switch {
	case kind == kindNil, kind == declared:
		return v, true
	case declared == kindDram && kind == kindGrain:
		return float64(v.(int64)), true
	}

	return nil, false
}
package scanner

type Scanner struct {
	source  string
	start   int
	current int
	line    int
}

// Constructor
func NewScanner(source string) *Scanner {
	return &Scanner{
		source: source,
		line:   1,
	}
}

// Main scanning function
func (s *Scanner) ScanTokens() []Token {
	tokens := []Token{}

	//TODO: implement scanning process

	return tokens
}

func (s *Scanner) advance() byte {
	char := s.source[s.current]
	s.current++
	return char
}

func (s *Scanner) peek() byte {
	if s.current >= len(s.source) {
		return '\x00'
	}
	return s.source[s.current]
}

func (s *Scanner) match(expected byte) bool {
	if s.current >= len(s.source) {
		return false
	}

	if s.source[s.current] != expected {
		return false
	}

	s.current++
	return true
}

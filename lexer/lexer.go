package zeus_lexer

const (
	Word       = iota
	Puntuation = iota
	Quotation  = iota
	Unknown    = iota
)

type Lexer struct {
	Input    string
	Position int
}

func NewLexer(input string) *Lexer {
	return &Lexer{Input: input}
}

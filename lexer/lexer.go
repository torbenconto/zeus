package zeus_lexer

const (
	Word             = iota
	Punctuation      = iota
	Quotation        = iota
	Apostrophe       = iota
	Comma            = iota
	Number           = iota
	RightParenthesis = iota
	LeftParenthesis  = iota
	Symbol           = iota
	ForwardSlash     = iota
	Colon            = iota
	Unknown          = iota
)

type Lexer struct {
	Input    string
	Position int
}

func NewLexer(input string) *Lexer {
	return &Lexer{Input: input}
}

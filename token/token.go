package token

type Token struct {
	Type    TokenType
	Literal string
}

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
	// Identifiers + Lietrals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456
	// Operators
	ASSIGN = "="
	PLUS   = "+"
	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

/**
* Checks the keywords table to see whether the given identifier is in
* fact a keyword. If it is, it returns the keyword's TokenType constant. If it isn't, we
* just get back token. IDENT, which is the TokenType for all user-defined identifiers.
 */
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

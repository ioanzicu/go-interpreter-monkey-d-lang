package token

const (
	ILLEGAL = "ILLEGAL" // token not supported
	EOF     = "EOF"     // End Of File

	// Identifiers + literals
	IDENT  = "IDENT" // add, foobar, x, y ...
	INT    = "INT"   // 123456789
	STRING = "STRING"

	// Operators
	ASSIGN   = "="
	EQ       = "=="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	NOT_EQ   = "!="
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
	Error   error
}

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if ttype, ok := keywords[ident]; ok {
		return ttype
	}

	return IDENT
}

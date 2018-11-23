// Package token encodes tokens for the Monkey language.
package token

// Type represents the type of a Monkey token.
type Type string

// Token represents a single concrete Monkey token.
type Token struct {
	Type
	Literal string
}

// Enumeration of Types.
const (
	ILLEGAL Type = "ILLEGAL"
	EOF     Type = "EOF"

	// Identifiers and literals
	IDENT Type = "IDENT"
	INT   Type = "INT"

	// Operators
	ASSIGN   Type = "="
	PLUS     Type = "+"
	MINUS    Type = "-"
	BANG     Type = "!"
	ASTERISK Type = "*"
	SLASH    Type = "/"

	LT    Type = "<"
	GT    Type = ">"
	EQ    Type = "=="
	NOTEQ Type = "!="

	// Delimiters
	COMMA     Type = ","
	SEMICOLON Type = ";"
	LPAREN    Type = "("
	RPAREN    Type = ")"
	LBRACE    Type = "{"
	RBRACE    Type = "}"

	// Keywords
	FUNCTION Type = "FUNCTION"
	LET      Type = "LET"
	TRUE     Type = "TRUE"
	FALSE    Type = "FALSE"
	IF       Type = "IF"
	ELSE     Type = "ELSE"
	RETURN   Type = "RETURN"
)

var keywords = map[string]Type{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// LookupIdentifier returns a keyword or IDENT Type for a character string.
func LookupIdentifier(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

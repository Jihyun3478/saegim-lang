package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	IDENT = "IDENT"
	INT = "INT"

	ASSIGN = "="
	PLUS = "PLUS"
	MINUS = "MINUS"
	ASTERISK = "ASTERISK"
	SLASH = "SLASH"

	LT = "<"
	GT = ">"
	EQ = "=="
	NOT_EQ = "!="

	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	
	VARIABLE = "변수"
)

var keywords = map[string]TokenType{
	"변수": VARIABLE,
}

func CheckKeyword(keyword string) TokenType {
	if tok, ok := keywords[keyword]; ok {
		return tok
	}
	return IDENT
}

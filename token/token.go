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
	LBRACE = "{"
	RBRACE = "}"
	
	VARIABLE = "변수"
	IF = "만약"
	ELSE = "아니면"
	TRUE = "참"
	FALSE = "거짓"
)

var keywords = map[string]TokenType{
	"변수": VARIABLE,
	"만약": IF,
	"아니면": ELSE,
	"참": TRUE,
	"거짓": FALSE,
}

func CheckKeyword(keyword string) TokenType {
	if tok, ok := keywords[keyword]; ok {
		return tok
	}
	return IDENT
}

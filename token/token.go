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
	PLUS = "+"
	MINUS = "-"
	ASTERISK = "*"
	SLASH = "/"

	LT = "<"
	GT = ">"
	LTE = "<="
	GTE = ">="
	EQ = "=="
	NOT_EQ = "!="

	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	
	VARIABLE = "VARIABLE"
	IF = "IF"
	ELSE = "ELSE"
	TRUE = "TRUE"
	FALSE = "FALSE"
	FUNCTION = "FUNCTION"
	RETURN = "RETURN"
	STRING = "STRING"
)

type KeywordSet map[string]TokenType

func CheckKeyword(keyword string, keywordSet KeywordSet) TokenType {
	if tok, ok := keywordSet[keyword]; ok {
		return tok
	}
	return IDENT
}

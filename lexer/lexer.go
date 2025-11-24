package lexer

import (
	"github.com/Jihyun3478/saegim-lang/token"
)

type Lexer struct {
	input []rune
	position int
	readPosition int
	character rune
	keywordSet token.KeywordSet
}

func New(input string, keywordSet ...token.KeywordSet) *Lexer {
	ks := token.StandardKeywords
	if len(keywordSet) > 0 {
		ks = keywordSet[0]
	}

	lexer := &Lexer{
		input: []rune(input),
		keywordSet: ks,
	}
	lexer.readChar()
	return lexer
}

func (lexer *Lexer) NextToken() token.Token {
	var tok token.Token

	lexer.skipWhitespace()

	switch lexer.character {
	case '=':
		if lexer.peekChar() == '=' {
			ch := lexer.character
			lexer.readChar()
			tok = token.Token{
				Type: token.EQ,
				Literal: string(ch) + string(lexer.character),
			}
		} else {
			tok = newToken(token.ASSIGN, lexer.character)
		}
	case '+':
		tok = newToken(token.PLUS, lexer.character)
	case '-':
		tok = newToken(token.MINUS, lexer.character)
	case '*':
		tok = newToken(token.ASTERISK, lexer.character)
	case '/':
		tok = newToken(token.SLASH, lexer.character)
	case '<':
		if lexer.peekChar() == '=' {
			ch := lexer.character
			lexer.readChar()
			tok = token.Token{
				Type: token.LTE,
				Literal: string(ch) + string(lexer.character),
			}
		} else {
			tok = newToken(token.LT, lexer.character)
		}
	case '>':
		if lexer.peekChar() == '=' {
			ch := lexer.character
			lexer.readChar()
			tok = token.Token{
				Type: token.GTE,
				Literal: string(ch) + string(lexer.character),
			}
		} else {
			tok = newToken(token.GT, lexer.character)
		}
	case '!':
		if lexer.peekChar() == '=' {
			ch := lexer.character
			lexer.readChar()
			tok = token.Token{
				Type: token.NOT_EQ,
				Literal: string(ch) + string(lexer.character),
			}
		} else {
			tok = newToken(token.ILLEGAL, lexer.character)
		}
	case '(':
		tok = newToken(token.LPAREN, lexer.character)
	case ')':
		tok = newToken(token.RPAREN, lexer.character)
	case '{':
		tok = newToken(token.LBRACE, lexer.character)
	case '}':
		tok = newToken(token.RBRACE, lexer.character)
	case ',':
		tok = newToken(token.COMMA, lexer.character)
	case ';':
		tok = newToken(token.SEMICOLON, lexer.character)
	case '"':
		tok.Type = token.STRING
		tok.Literal = lexer.readString()
		return tok
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isKorean(lexer.character) {
			tok.Literal = lexer.readIdentifier()
			tok.Type = token.CheckKeyword(tok.Literal, lexer.keywordSet)
			return tok
		} else if isDigit(lexer.character) {
			tok.Type = token.INT
			tok.Literal = lexer.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, lexer.character)
		}
	
	}

	lexer.readChar()
	return tok
}

func (lexer *Lexer) readChar() {
	if lexer.readPosition >= len(lexer.input) {
		lexer.character = 0
	} else {
		lexer.character = lexer.input[lexer.readPosition]
	}
	lexer.position = lexer.readPosition
	lexer.readPosition += 1
}

func (lexer *Lexer) peekChar() rune {
	if lexer.readPosition >= len(lexer.input) {
		return 0
	} else {
		return lexer.input[lexer.readPosition]
	}
}

func newToken(tokenType token.TokenType, character rune) token.Token {
	return token.Token{Type: tokenType, Literal: string(character)}
}

func isKorean(character rune) bool {
	return 0xAC00 <= character && character <= 0xD7A3
}

func isDigit(character rune) bool {
	return '0' <= character && character <= '9'
}

func (lexer *Lexer) readIdentifier() string {
	position := lexer.position
	for isKorean(lexer.character) {
		lexer.readChar()
	}
	return string(lexer.input[position:lexer.position])
}

func (lexer *Lexer) readNumber() string {
	position := lexer.position
	for isDigit(lexer.character) {
		lexer.readChar()
	}
	return string(lexer.input[position:lexer.position])
}

func (lexer *Lexer) readString() string {
	position := lexer.position + 1
	for {
		lexer.readChar()
		if lexer.character == '"' || lexer.character == 0 {
			break
		}
	}
	str := string(lexer.input[position:lexer.position])
	lexer.readChar()
	
	return str
}

func (lexer *Lexer) skipWhitespace() {
	for lexer.character == ' ' || lexer.character == '\t' || lexer.character == '\n' || lexer.character == '\r' {
		lexer.readChar()
	}
}

package parser

import (
	"fmt"
	"strconv"

	"github.com/Jihyun3478/saegim-lang/token"
	"github.com/Jihyun3478/saegim-lang/lexer"
	"github.com/Jihyun3478/saegim-lang/ast"
)

type Parser struct {
	lexer *lexer.Lexer
	errors []string

	currentToken token.Token
	peekToken token.Token
}

func New(lexer *lexer.Lexer) *Parser {
	parser := &Parser{
		lexer: lexer,
		errors: []string{},
	}

	parser.nextToken()
	parser.nextToken()

	return parser
}

func (parser *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for parser.currentToken.Type != token.EOF {
		statement := parser.parseStatement()
		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}
		parser.nextToken()
	}

	return program
}

func (parser *Parser) nextToken() {
	parser.currentToken = parser.peekToken
	parser.peekToken = parser.lexer.NextToken()
}

func (parser *Parser) parseStatement() ast.Statement {
	switch parser.currentToken.Type {
	case token.VARIABLE:
		return parser.parseVariableStatement()
	default:
		return nil
	}
}

func (parser *Parser) parseVariableStatement() *ast.VariableStatement {
	statement := &ast.VariableStatement{Token: parser.currentToken}

	if !parser.expectPeek(token.IDENT) {
		return nil
	}

	statement.Name = &ast.Identifier{
		Token: parser.currentToken,
		Value: parser.currentToken.Literal,
	}

	if !parser.expectPeek(token.ASSIGN) {
		return nil
	}

	parser.nextToken()

	statement.Value = parser.parseIntegerLiteral()

	if parser.peekTokenIs(token.SEMICOLON) {
		parser.nextToken()
	}

	return statement
}

func (parser *Parser) parseIntegerLiteral() ast.Expression {
	integerLiteral := &ast.IntegerLiteral{Token: parser.currentToken}

	value, err := strconv.ParseInt(parser.currentToken.Literal, 0, 64)
	if err != nil {
		message := fmt.Sprintf("could not parse %q as integer", parser.currentToken.Literal)
		parser.errors = append(parser.errors, message)
		return nil
	}

	integerLiteral.Value = value

	return integerLiteral
}

func (parser *Parser) peekTokenIs(t token.TokenType) bool {
	return parser.peekToken.Type == t
}

func (parser *Parser) expectPeek(t token.TokenType) bool {
	if parser.peekTokenIs(t) {
		parser.nextToken()
		return true
	} else {
		parser.peekError(t)
		return false
	}
}

func (parser *Parser) peekError(t token.TokenType) {
	message := fmt.Sprintf("expected next token to be %s, got %s instead", t, parser.peekToken.Type)
	parser.errors = append(parser.errors, message)
}

func (parser *Parser) Errors() []string {
	return parser.errors
}

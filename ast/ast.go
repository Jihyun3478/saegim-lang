package ast

import (
	"bytes"
	"strings"

	"github.com/Jihyun3478/saegim-lang/token"
)

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

type VariableStatement struct {
	Token token.Token
	Name *Identifier
	Value Expression
}

func (variable *VariableStatement) statementNode() {}

func (variable *VariableStatement) TokenLiteral() string {
	return variable.Token.Literal
}

func (variable *VariableStatement) String() string {
	var out bytes.Buffer

	out.WriteString(variable.Token.Literal + " ")
	out.WriteString(variable.Name.String())
	out.WriteString(" = ")

	if variable.Value != nil {
		out.WriteString(variable.Value.String())
	}

	out.WriteString(";")
	return out.String()
}

type ExpressionStatement struct {
	Token token.Token
	Expression Expression
}

func (expression *ExpressionStatement) statementNode() {}

func (expression *ExpressionStatement) TokenLiteral() string {
	return expression.Token.Literal
}

func (expression *ExpressionStatement) String() string {
	if expression.Expression != nil {
		return expression.Expression.String()
	}
	return ""
}

type BlockStatement struct {
	Token token.Token
	Statements []Statement
}

func (block *BlockStatement) statementNode() {}

func (block *BlockStatement) TokenLiteral() string {
	return block.Token.Literal
}

func (block *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range block.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

type Identifier struct {
	Token token.Token
	Value string
}

func (identifier *Identifier) expressionNode() {}

func (identifier *Identifier) TokenLiteral() string {
	return identifier.Token.Literal
}

func (identifier *Identifier) String() string {
	return identifier.Value
}

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (integerLiteral *IntegerLiteral) expressionNode() {}

func (integerLiteral *IntegerLiteral) TokenLiteral() string {
	return integerLiteral.Token.Literal
}

func (integerLiteral *IntegerLiteral) String() string {
	return integerLiteral.Token.Literal
}

type InfixExpression struct {
	Token token.Token
	Left Expression
	Operator string
	Right Expression
}

func (infix *InfixExpression) expressionNode() {}

func (infix *InfixExpression) TokenLiteral() string {
	return infix.Token.Literal
}

func (infix *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(infix.Left.String())
	out.WriteString(" " + infix.Operator + " ")
	out.WriteString(infix.Right.String())
	out.WriteString(")")

	return out.String()
}

type IfExpression struct {
	Token token.Token
	Condition Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (ifExpression *IfExpression) expressionNode() {}

func (ifExpression *IfExpression) TokenLiteral() string {
	return ifExpression.Token.Literal
}

func (ifExpression *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("만약")
	out.WriteString(ifExpression.Condition.String())
	out.WriteString(" ")
	out.WriteString(ifExpression.Consequence.String())

	if ifExpression.Alternative != nil {
		out.WriteString("아니면 ")
		out.WriteString(ifExpression.Alternative.String())
	}

	return out.String()
}

type Boolean struct {
	Token token.Token
	Value bool
}

func (boolean *Boolean) expressionNode() {}

func (boolean *Boolean) TokenLiteral() string {
	return boolean.Token.Literal
}

func (boolean *Boolean) String() string {
	return boolean.Token.Literal
}

type FunctionalLiteral struct {
	Token token.Token
	Parameters []*Identifier
	Body *BlockStatement
}

func (functionalLiteral *FunctionalLiteral) expressionNode() {}

func (functionalLiteral *FunctionalLiteral) TokenLiteral() string {
	return functionalLiteral.Token.Literal
}

func (functionalLiteral *FunctionalLiteral) String() string {
	var out bytes.Buffer

	parameters := []string{}
	for _, p := range functionalLiteral.Parameters {
		parameters = append(parameters, p.String())
	}

	out.WriteString(functionalLiteral.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(parameters, ", "))
	out.WriteString(") ")
	out.WriteString(functionalLiteral.Body.String())

	return out.String()
}

type CallExpression struct {
	Token token.Token
	Function Expression
	Arguments []Expression
}

func (callExpression *CallExpression) expressionNode() {}

func (callExpression *CallExpression) TokenLiteral() string {
	return callExpression.Token.Literal
}

func (callExpression *CallExpression) String() string {
	var out bytes.Buffer

	arguments := []string{}
	for _, a := range callExpression.Arguments {
		arguments = append(arguments, a.String())
	}

	out.WriteString(callExpression.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(arguments, ", "))
	out.WriteString(") ")

	return out.String()
}

type ReturnStatement struct {
	Token token.Token
	ReturnValue Expression
}

func (returnStatement *ReturnStatement) statementNode() {}

func (returnStatement *ReturnStatement) TokenLiteral() string {
	return returnStatement.Token.Literal
}

func (returnStatement *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(returnStatement.TokenLiteral() + " ")

	if returnStatement.ReturnValue != nil {
		out.WriteString(returnStatement.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

// Package ast defines an Abstract Syntax Tree for the Monkey programming
// language.
package ast

import (
	"bytes"

	"github.com/pto/monkey/token"
)

// Node is an AST node.
type Node interface {
	TokenLiteral() string
	String() string
}

// Statement is a Node for statements.
type Statement interface {
	Node
	statementNode() // type marker only
}

// Expression is a Node for expressions.
type Expression interface {
	Node
	expressionNode() // type market only
}

// Program is a Node that represents an entire program.
type Program struct {
	Statements []Statement
}

// TokenLiteral returns the first statement in a program, for debugging.
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// String returns a description of the Program.
func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

// LetStatement is a Node representing a let statement.
type LetStatement struct {
	Token token.Token // always a LET
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

// TokenLiteral for a let statement always returns "let".
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// String returns a description of the LetStatement.
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")
	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")

	return out.String()
}

// Identifier is a Node representing an identifier.
type Identifier struct {
	Token token.Token // always IDENT
	Value string
}

func (i *Identifier) expressionNode() {}

// TokenLiteral for an identifier returns the name of the identifier.
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

// String returns a description of the Identifier.
func (i *Identifier) String() string {
	return i.Value
}

// ReturnStatement is a Node representing a return statement.
type ReturnStatement struct {
	Token       token.Token // always a RETURN
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

// TokenLiteral for a return statement always returns "return".
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

// String returns a description of the ReturnStatement.
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")

	return out.String()
}

// ExpressionStatement is a Node representing an expression statement.
type ExpressionStatement struct {
	Token      token.Token // the first token only
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

// TokenLiteral for an expression statement returns the first token literal.
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

// String returns a description of the ExpressionStatement.
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

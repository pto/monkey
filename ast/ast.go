// Package ast defines an Abstract Syntax Tree for the Monkey programming
// language.
package ast

import "github.com/pto/monkey/token"

// Node is an AST node.
type Node interface {
	TokenLiteral() string
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

// ReturnStatement is a Node repsenting a return statement.
type ReturnStatement struct {
	Token       token.Token // always a RETURN
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

// TokenLiteral for a return statement always returns "return".
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

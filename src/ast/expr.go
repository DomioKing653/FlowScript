package ast

import "github.com/DomioKing653/FlowScript/src/lexer"

//Literal

type NumberExpr struct {
	Value float64
}

func (expr *NumberExpr) expr() {

}

type StringExpr struct {
	Value string
}

func (expr StringExpr) expr() {

}

type SymbolExpr struct {
	Value string
}

func (expr SymbolExpr) expr() {

}

//Complex

type BinaryOperation struct {
	Left     Expression
	Operator lexer.Token
	Right    Expression
}

func (bin BinaryOperation) expr() {

}

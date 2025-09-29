package ast

import "github.com/DomioKing653/FlowScript/src/lexer"

//Literal

type NumberExpr struct {
	value float64
}

func (expr *NumberExpr) expr() {

}

type StringExpr struct {
	value float64
}

func (expr StringExpr) expr() {

}

type SymbolExpr struct {
	value float64
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

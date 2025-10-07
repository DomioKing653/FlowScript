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

type PrefixEpr struct {
	Operator  lexer.Token
	RightExpr Expression
}

func (n PrefixEpr) expr() {

}

type AssigmentEpr struct {
	AssignedVar Expression
	Operator    lexer.Token
	Value       Expression
}

func (n AssigmentEpr) expr() {

}

type StructInstantiation struct {
	StructName       string
	StructPropreties map[string]Expression
}

func (n StructInstantiation) expr() {

}

type ArrayInstantiation struct {
	Lenght     lexer.Token
	Underlying Type
	Contents   []Expression
}

func (n ArrayInstantiation) expr() {

}

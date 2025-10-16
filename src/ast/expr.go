package ast

import (
	"github.com/DomioKing653/FlowScript/src/lexer"
	"github.com/DomioKing653/FlowScript/src/runtime"
)

//Literal

type NumberExpr struct {
	Value float64
}

func (expr *NumberExpr) expr() runtime.RuntimeValue {
	return runtime.RuntimeValue{
		NumberVal: expr.Value,
	}
}

type StringExpr struct {
	Value string
}

func (expr StringExpr) expr() runtime.RuntimeValue {
	return runtime.RuntimeValue{
		StringVal: expr.Value,
	}
}

type SymbolExpr struct {
	Value string
}

func (expr SymbolExpr) expr() runtime.RuntimeValue {
	return runtime.RuntimeValue{
		StringVal: expr.Value,
	}
}

//Complex

type BinaryOperation struct {
	Left     Expression
	Operator lexer.Token
	Right    Expression
}

func (bin BinaryOperation) expr() runtime.RuntimeValue {
	panic("not implemented")
}

type PrefixEpr struct {
	Operator  lexer.Token
	RightExpr Expression
}

func (n PrefixEpr) expr() runtime.RuntimeValue {
	panic("not implemented")
}

type AssigmentEpr struct {
	AssignedVar Expression
	Operator    lexer.Token
	Value       Expression
}

func (n AssigmentEpr) expr() runtime.RuntimeValue {
	panic("not implemented")
}

type StructInstantiation struct {
	StructName       string
	StructPropreties map[string]Expression
}

func (n StructInstantiation) expr() runtime.RuntimeValue {
	panic("not implemented")
}

type ArrayInstantiation struct {
	Lenght     lexer.Token
	Underlying Type
	Contents   []Expression
}

func (n ArrayInstantiation) expr() runtime.RuntimeValue {
	panic("not implemented")
}

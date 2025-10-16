package ast

import "github.com/DomioKing653/FlowScript/src/runtime"

type BlockStatment struct {
	Body []Statement
}

func (stmt BlockStatment) stmt() {

}

type ExprStatment struct {
	Expr Expression
}

func (expStmt ExprStatment) stmt() runtime.RuntimeValue {
	return expStmt.Expr.expr()
}

// Variables
type VarDeclStatment struct {
	VariableName string
	IsConst      bool
	Value        Expression
	ExplicitType Type
}

func (expStmt VarDeclStatment) stmt() runtime.RuntimeValue {
	panic("not implemented")
}

//Structs

type StructProprety struct {
	Type     Type
	IsStatic bool
}

type StructDeclStmt struct {
	StructName       string
	StructPropreties map[string]StructProprety
}

func (n StructDeclStmt) stmt() runtime.RuntimeValue {
	panic("not implemented")
}

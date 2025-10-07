package ast

type BlockStatment struct {
	Body []Statement
}

func (stmt BlockStatment) stmt() {

}

type ExprStatment struct {
	Expr Expression
}

func (expStmt ExprStatment) stmt() {

}

// Variables
type VarDeclStatment struct {
	VariableName string
	IsConst      bool
	Value        Expression
	ExplicitType Type
}

func (expStmt VarDeclStatment) stmt() {

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

func (n StructDeclStmt) stmt() {

}

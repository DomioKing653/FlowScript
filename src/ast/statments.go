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

type VarDeclStatment struct {
	VariableName string
	IsConst      bool
	Value        Expression
}

func (expStmt VarDeclStatment) stmt() {

}

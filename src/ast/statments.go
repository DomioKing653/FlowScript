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

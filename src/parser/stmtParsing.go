package parser

import (
	"github.com/DomioKing653/FlowScript/src/ast"
	"github.com/DomioKing653/FlowScript/src/lexer"
)

func parse_stmt(p *parser) ast.Statement {
	stmt_fn, exists := stmt_lu[p.currentTokenKind()]
	if exists {
		return stmt_fn(p)
	}
	expression := parse_expr(p, default_bp)
	p.expect(lexer.SEMI_COLON)
	return ast.ExprStatment{Expr: expression}

}

func parse_var_decl_statment(p *parser) ast.Statement {
	isConst := p.advance().Kind == lexer.CONST
	varName := p.expectError(lexer.IDENTIFIER, "Error::Syntax->Expected identifier in variable declaration").Value
	p.expect(lexer.ASSIGNMENT)
	varValue := parse_expr(p, assigment)
	p.expect(lexer.SEMI_COLON)
	return ast.VarDeclStatment{IsConst: isConst, VariableName: varName, Value: varValue}
}

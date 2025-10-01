package parser

import (
	"github.com/DomioKing653/FlowScript/src/ast"
	"github.com/DomioKing653/FlowScript/src/lexer"
)

func parse_stmt(p *parser) ast.Statement {
	stmt_fn, exists := stmt_lu[p.currentTokenKind()]
	if exists {
		return stmt_fn(p)
	} else {
		expression := parse_expr(p, default_bp)
		p.expect(lexer.SEMI_COLON)
		return ast.ExprStatment{Expr: expression}
	}
}

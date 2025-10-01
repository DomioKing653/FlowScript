package parser

import (
	"fmt"
	"strconv"

	"github.com/DomioKing653/FlowScript/src/ast"
	"github.com/DomioKing653/FlowScript/src/lexer"
)

func parse_primary_expr(p *parser) ast.Expression {
	switch p.currentTokenKind() {
	case lexer.NUMBER:
		number, _ := strconv.ParseFloat(p.advance().Value, 64)
		return &ast.NumberExpr{Value: number}
	case lexer.STRING:
		return &ast.StringExpr{Value: p.advance().Value}
	case lexer.IDENTIFIER:
		return &ast.SymbolExpr{Value: p.advance().Value}
	default:
		panic(fmt.Sprintf("cannot create primary_expr from %s\n", lexer.TokenKindString(p.currentTokenKind())))

	}
}

package parser

import (
	"fmt"
	"strconv"

	"github.com/DomioKing653/FlowScript/src/ast"
	"github.com/DomioKing653/FlowScript/src/lexer"
)

func parse_expr(p *parser, bp binding_power) ast.Expression {
	tokenKind := p.currentTokenKind()
	nud_fn, exists := nud_lu[tokenKind]

	if !exists {
		panic(fmt.Sprintf("Error::Parsing->NUD HANDLER EXPECTED TOKEN %s\n", lexer.TokenKindString(tokenKind)))
	}
	left := nud_fn(p)

	for bp_lu[p.currentTokenKind()] > bp {
		tokenKind = p.currentTokenKind()
		led_fn, exists := led_lu[tokenKind]
		if !exists {
			panic(fmt.Sprintf("Error::Parsing->LED HANDLER EXPECTED TOKEN %s\n", lexer.TokenKindString(tokenKind)))
		}
		left = led_fn(p, left, bp)
	}
	return left
}

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
		panic(fmt.Sprintf("Error::Parser->cannot create primary_expr from %s\n", lexer.TokenKindString(p.currentTokenKind())))
	}
}

func parse_binary_expr(p *parser, left ast.Expression, bp binding_power) ast.Expression {
	operrator := p.advance()
	right := parse_expr(p, bp)
	return &ast.BinaryOperation{
		Operator: operrator,
		Left:     left,
		Right:    right,
	}
}

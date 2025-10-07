package parser

import (
	"fmt"
	"strconv"

	"github.com/DomioKing653/FlowScript/src/ast"
	"github.com/DomioKing653/FlowScript/src/helpers"
	"github.com/DomioKing653/FlowScript/src/lexer"
)

func parse_expr(p *parser, bp binding_power) ast.Expression {
	tokenKind := p.currentTokenKind()
	nud_fn, exists := nud_lu[tokenKind]

	if !exists {
		panic(fmt.Sprintf("Error::Parsing->No NUD handler for token %s", lexer.TokenKindString(tokenKind)))
	}
	left := nud_fn(p)
	for bp_lu[p.currentTokenKind()] > bp {
		tokenKind = p.currentTokenKind()
		led_fn, exists := led_lu[tokenKind]
		if !exists {
			panic(fmt.Sprintf("Error::Parsing->No LED handler for token %s", lexer.TokenKindString(tokenKind)))
		}
		left = led_fn(p, left, bp_lu[p.currentTokenKind()])
	}
	return left
}

func parse_primary_expr(p *parser) ast.Expression {
	switch p.currentTokenKind() {
	case lexer.NUMBER:
		number, _ := strconv.ParseFloat(p.advance().Value, 64)
		return &ast.NumberExpr{
			Value: number,
		}
	case lexer.STRING:
		return &ast.StringExpr{
			Value: p.advance().Value,
		}
	case lexer.IDENTIFIER:
		return &ast.SymbolExpr{
			Value: p.advance().Value,
		}
	default:
		panic(fmt.Sprintf("Error::Parsing->Cannot create primary expression from token %s", lexer.TokenKindString(p.currentTokenKind())))
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

func parse_assigment_expr(p *parser, left ast.Expression, bp binding_power) ast.Expression {
	token := p.advance()
	rhs := parse_expr(p, assigment)
	return ast.AssigmentEpr{
		Value:       rhs,
		Operator:    token,
		AssignedVar: left,
	}
}

func parse_prefix_expression(p *parser) ast.Expression {
	operatorToken := p.advance()
	rhs := parse_expr(p, default_bp)
	return ast.PrefixEpr{
		Operator:  operatorToken,
		RightExpr: rhs,
	}
}

func parse_grouping_expr(p *parser) ast.Expression {
	p.advance()
	expr := parse_expr(p, default_bp)
	p.expect(lexer.CLOSE_PAREN)
	return expr
}

func parse_struct_instantiation_expr(p *parser, left ast.Expression, bp binding_power) ast.Expression {
	var structName = helpers.ExpectedType[ast.SymbolExpr](left).Value
	var structPropreties = map[string]ast.Expression{}
	p.expect(lexer.OPEN_CURLY)
	for p.hasTokens() && p.currentTokenKind() != lexer.CLOSE_CURLY {
		propretyName := p.expect(lexer.IDENTIFIER).Value
		p.expect(lexer.COLON)
		value := parse_expr(p, logical)
		
		structPropreties[propretyName] = value
		if p.currentTokenKind() != lexer.COMMA {
			panic(fmt.Sprintf("Error::Synax->Expected comma in %s instantiation", structName))
		}
		p.advance()
	}
	p.expect(lexer.CLOSE_CURLY)
	return ast.StructInstantiation{
		StructName:       structName,
		StructPropreties: structPropreties,
	}
}

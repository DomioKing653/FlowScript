package parser

import (
	"github.com/DomioKing653/FlowScript/src/ast"
	"github.com/DomioKing653/FlowScript/src/lexer"
)

type type_nud_handler func(p *parser) ast.Type
type type_led_handler func(p *parser, left ast.Type, bp binding_power) ast.Type

type type_nud_lookup map[lexer.TokenKind]type_nud_handler
type type_led_lookup map[lexer.TokenKind]type_led_handler
type type_bp_lookup map[lexer.TokenKind]binding_power

var type_bp_lu = bp_lookup{}
var type_nud_lu = type_nud_lookup{}
var type_led_lu = type_led_lookup{}

func type_led(kind lexer.TokenKind, bp binding_power, led_fn type_led_handler) {
	type_bp_lu[kind] = bp
	type_led_lu[kind] = led_fn
}

func type_nud(kind lexer.TokenKind, nud_fn type_nud_handler) {
	type_nud_lu[kind] = nud_fn
}

func CreateTokenTypeLookups() {
	type_nud(lexer.IDENTIFIER, parse_symbol_type)
}

func parse_symbol_type(p *parser) ast.Type {
	return &ast.SymobolType{
		Name: p.expect(lexer.IDENTIFIER).Value,
	}
}

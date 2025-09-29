package parser

import (
	"github.com/DomioKing653/FlowScript/src/ast"

	"github.com/DomioKing653/FlowScript/src/lexer"
)

type parser struct {
	tokens []lexer.Token
	pos    int
}

func Parse(tokens lexer.Token) ast.BlockStatment {
	Body := make([]ast.Statement, 0)
	p := CreateParser([]lexer.Token{})
	for p.hasTokens() {
		Body = append(Body, parse_stmt(p))
	}
	return ast.BlockStatment{Body: Body}
}

// HELPER METHODS

func CreateParser(tokens []lexer.Token) *parser {
	return &parser{
		tokens: tokens, pos: 0,
	}
}

func (p *parser) currentToken() lexer.Token {
	return p.tokens[p.pos]
}

func (p *parser) currentTokenKind() lexer.TokenKind {
	return p.currentToken().Kind
}

func (p *parser) advance() lexer.Token {
	tk := p.currentToken()
	p.pos++
	return tk
}

func (p *parser) hasTokens() bool {
	return p.pos > len(p.tokens) && p.currentTokenKind() != lexer.EOF
}

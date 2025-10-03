package parser

import (
	"fmt"

	"github.com/DomioKing653/FlowScript/src/ast"

	"github.com/DomioKing653/FlowScript/src/lexer"
)

type parser struct {
	tokens []lexer.Token
	pos    int
}

func Parse(tokens []lexer.Token) ast.BlockStatment {
	Body := make([]ast.Statement, 0)
	p := CreateParser(tokens)
	for p.hasTokens() {
		Body = append(Body, parse_stmt(p))
	}
	return ast.BlockStatment{Body: Body}
}

// HELPER METHODS

func CreateParser(tokens []lexer.Token) *parser {
	CreateLookups()
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
	return p.pos < len(p.tokens) && p.currentTokenKind() != lexer.EOF
}

func (p *parser) expectError(expectedKind lexer.TokenKind, err any) lexer.Token {
	token := p.currentToken()
	kind := token.Kind

	if kind != expectedKind {
		if err == nil {
			err = fmt.Sprintf("Error::Syntax->Expected %s but received %s", lexer.TokenKindString(expectedKind), lexer.TokenKindString(kind))
		}

		panic(err)
	}

	return p.advance()
}

func (p *parser) expect(expectedKind lexer.TokenKind) lexer.Token {
	return p.expectError(expectedKind, nil)
}

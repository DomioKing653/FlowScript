package parser

import (
	"fmt"

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
	var explicitType ast.Type
	var varValue ast.Expression
	isConst := p.advance().Kind == lexer.CONST
	varName := p.expectError(lexer.IDENTIFIER, "Error::Syntax->Expected identifier in variable declaration").Value
	if p.currentTokenKind() == lexer.COLON {
		p.advance() //Eat colon
		explicitType = parse_type(p, default_bp)
	}
	if p.currentTokenKind() != lexer.SEMI_COLON {
		p.expect(lexer.ASSIGNMENT)
		varValue = parse_expr(p, assigment)
	} else if explicitType == nil {
		panic("Error::Parser->Missing either assigment or explicit type")
	}

	if isConst && varValue == nil {
		panic("Error::Compile time->Can't have const without a value")
	}

	p.expect(lexer.SEMI_COLON)
	return ast.VarDeclStatment{
		IsConst:      isConst,
		VariableName: varName,
		Value:        varValue,
		ExplicitType: explicitType,
	}
}

func parse_struct_decl_statment(p *parser) ast.Statement {
	p.expect(lexer.STRUCT)
	var structName = p.expect(lexer.IDENTIFIER).Value
	var structPropreties = map[string]ast.StructProprety{}
	p.expect(lexer.OPEN_CURLY)

	for p.hasTokens() && p.currentToken().Kind != lexer.CLOSE_CURLY {
		var isStatic bool
		var propretyName string
		var propretyType ast.Type
		if p.currentTokenKind() == lexer.STATIC {
			isStatic = true
			p.expect(lexer.STATIC)
		}

		if p.currentTokenKind() == lexer.IDENTIFIER {
			propretyName = p.advance().Value
			p.expectError(lexer.COLON, "Erros::Syntax->Expected to find colon in struct definition")
			propretyType = parse_type(p, default_bp)
			p.expect(lexer.SEMI_COLON)

			_, exists := structPropreties[propretyName]
			if exists {
				panic(fmt.Sprintf("Error::Compiletime->Proprety '%s' alredy exists for struct '%s'", propretyName, structName))
			}
			structPropreties[propretyName] = ast.StructProprety{
				IsStatic: isStatic,
				Type:     propretyType,
			}
			continue
		}

	}

	p.expect(lexer.CLOSE_CURLY)
	return ast.StructDeclStmt{StructPropreties: structPropreties, StructName: structName}
}

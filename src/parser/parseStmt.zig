const std = @import("std");
const stmt = @import("../ast/statement.zig");
const parser = @import("parser.zig");
const Tokens = @import("../lexer/tokens.zig");
//Errors
pub const ParserErrors = error{ NotImplemented, ExpectedStatement, UnexpectedToken, UnexpectedEOF };
//Parsing
pub fn parseStmt(p: *parser.Parser) !stmt.Statement {
    switch (p.current_token.Kind) {
        Tokens.TokenKind.LET, Tokens.TokenKind.CONST => {
            const parsed_stmt = try parse_var_decl(p);
            _ = try p.expect(Tokens.TokenKind.SEMI_COLON);
            try p.advance();
            return parsed_stmt;
        },
        else => return ParserErrors.ExpectedStatement,
    }
    try p.advance();
}

fn parse_var_decl(p: *parser.Parser) !stmt.Statement {
    var is_const: bool = undefined;
    if (p.current_token.Kind == Tokens.TokenKind.CONST) {
        is_const = true;
    } else {
        is_const = false;
    }
    try p.advance();
    const id = try p.expect(Tokens.TokenKind.SYMBOL);
    try p.advance();
    if (p.current_token.Kind == Tokens.TokenKind.SEMI_COLON) {
        return stmt.Statement{ .varStmt = .{ .is_const = is_const, .id = id, .value = undefined } };
    } else {
        //const value = undefined;
    }
    const var_stmt = try p.alloc.create(stmt.varDeclStmt);

    var_stmt.* = .{ .id = id, .is_const = is_const, .value = undefined };
    _ = stmt.IStatement.init(var_stmt);
    return stmt.Statement{ .varStmt = .{ .is_const = is_const, .id = id, .value = undefined } };
}

fn parse_buildin() !stmt.Statement {}

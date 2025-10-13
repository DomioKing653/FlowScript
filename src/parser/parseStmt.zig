const std = @import("std");
const stmt = @import("../ast/statement.zig");
const parser = @import("parser.zig");
const Tokens = @import("../lexer/tokens.zig");
//Errors
pub const ParserErrors = error{ NotImplemented, ExpectedStatement, UnexpectedToken };
//Parsing
pub fn parseStmt(p: *parser.Parser) !stmt.Statement {
    switch (p.current_token.Kind) {
        Tokens.TokenKind.LET, Tokens.TokenKind.CONST => {
            const parsed_stmt = try parser_var_decl(p);
            _ = try p.expect(Tokens.TokenKind.SEMI_COLON);
            return parsed_stmt;
        },
        else => return ParserErrors.ExpectedStatement,
    }
    try p.advance();
}

fn parser_var_decl(p: *parser.Parser) !stmt.Statement {
    var is_const: bool = undefined;
    if (p.current_token.Kind == Tokens.TokenKind.CONST) {
        is_const = true;
    } else {
        is_const = true;
    }
    try p.advance();
    const id = try p.expect(Tokens.TokenKind.SYMBOL);
    try p.advance();
    if (p.current_token.Kind == Tokens.TokenKind.SEMI_COLON) {
        return stmt.Statement{ .varStmt = .{ .is_const = is_const, .id = id, .value = undefined } };
    } else {
        //const value = undefined;
    }
    return stmt.Statement{ .varStmt = .{ .is_const = is_const, .id = id, .value = undefined } };
}

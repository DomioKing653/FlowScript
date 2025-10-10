const stmt = @import("../ast/statement.zig");
const Tokens = @import("../lexer/tokens.zig");
const std = @import("std");
pub const Parser = struct {
    tokens: []Tokens.Token,
    pos_idx: usize,
    current_token: Tokens.Token,
    statements: stmt.Statement,
    pub fn parse(self: *Parser) !void {
        self.current_token = self.tokens[0];
        self.pos_idx += 1;
        while (self.current_token.Kind != Tokens.TokenKind.EOF) {
            try self.advance();
        }
        std.debug.print("Parsing ended", .{});
    }
    fn advance(self: *Parser) !void {
        self.current_token = self.tokens[self.pos_idx];
        self.pos_idx += 1;
    }
};

const stmt = @import("../ast/statement.zig");
const parseStmt = @import("parseStmt.zig");
const Tokens = @import("../lexer/tokens.zig");
const std = @import("std");
pub const Parser = struct {
    tokens: []Tokens.Token,
    pos_idx: usize,
    current_token: Tokens.Token,
    statements: std.ArrayList(stmt.Statement),
    alloc: std.mem.Allocator,
    pub fn parse(self: *Parser) !void {
        self.current_token = self.tokens[0];
        self.pos_idx += 1;
        while (self.current_token.Kind != Tokens.TokenKind.EOF) {
            try self.statements.append(self.alloc, try parseStmt.parseStmt(self));
        }
        std.debug.print("Parsing ended", .{});
    }
    pub fn advance(self: *Parser) !void {
        self.current_token = self.tokens[self.pos_idx];
        self.pos_idx += 1;
    }

    pub fn expect(self: *Parser, kind: Tokens.TokenKind) ![]u8 {
        if (self.current_token.Kind != kind) {
            return parseStmt.ParserErrors.UnexpectedToken;
        } else {
            return self.alloc.dupe(u8, self.current_token.Value);
        }
    }
};

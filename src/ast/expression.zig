const ast = @import("ast.zig");
const tokens = @import("../lexer/tokens.zig");
const std = @import("std");

pub const value = union(enum) { string_val: []u8, num_val: f64, bool_val: bool };

const BinaryExpr = struct {
    left: *ast.Expression,
    op: tokens.Token,
    right: *ast.Expression,

    pub fn expr(self: *BinaryExpr) !value {
        switch (self.op) {
            tokens.TokenKind.PLUS => value{ .num_val = self.left.expr() + self.right.expr() },
            tokens.TokenKind.PLUS => value{ .num_val = self.left.expr() - self.right.expr() },
            tokens.TokenKind.PLUS => value{ .num_val = self.left.expr() + self.right.expr() },
        }
    }
};
pub const Prefic_Expr = struct { left: tokens.Token, right: *ast.Expression };

const tokens = @import("../lexer/tokens.zig");
const expressions = @import("expression.zig");
const std = @import("std");

pub const Expression = union(enum) {
    number: f64,
    binary: expressions.BinaryExpr,
    pub fn expr(self: *Expression) !expressions.value {
        switch (self.*) {
            .binary => |*p| p.expr() catch |err| {
                std.fs.File.stdout().write(err);
            },
            .number => |*n| n.expr(),
        }
    }
};

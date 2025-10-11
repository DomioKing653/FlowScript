const std = @import("std");
const stmt = @import("../ast/statement.zig");
const parser = @import("parser.zig");
const ParserErrors = error{NotImplemented};
pub fn parseStmt(p: *parser.Parser) !void {
    std.debug.print("{any}\n", .{p.current_token.Kind});
    try p.advance();
}

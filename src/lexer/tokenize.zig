const std = @import("std");
const tokens = @import("tokens.zig");

pub fn tokenize(code: []const u8) !void {
    const stdout = std.io.getStdOut().writer();
    // writeAll writes raw bytes; change this later to actual tokenization
    try stdout.writeAll(code);
}

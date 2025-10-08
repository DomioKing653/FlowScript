const std = @import("std");
const Lexer = @import("lexer/tokenize.zig");

pub fn main() !void {
    const allocator = std.heap.page_allocator;
    const path = "C:/Users/simon/AAAProjects/AAAFlowScript/test.flw";

    const cwd = std.fs.cwd();
    const file = try cwd.openFile(path, .{});
    defer file.close();

    const bytes = try std.io.readAllAlloc(allocator, file.reader(), 0);
    defer allocator.free(bytes);

    try Lexer.tokenize(bytes);
}

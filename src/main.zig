const std = @import("std");
const Lexer = @import("lexer/tokenize.zig");

pub fn main() !void {
    const allocator = std.heap.page_allocator;
    const path = "C:/Users/simon/AAAProjects/AAAFlowScript/test.flw";

    // opening
    const file = try std.fs.cwd().openFile(path, .{});
    defer file.close();

    //reading file
    const bytes = try file.readToEndAlloc(allocator, std.math.maxInt(usize));
    defer allocator.free(bytes);

    const toks = try Lexer.tokenize(bytes);
    for (toks) |tok| {
        switch (tok.Value) {
            .char => |c| std.debug.print("Token: {c}\n", .{c}),
            .chars => |txt| std.debug.print("Token: {s}\n", .{txt}),
        }
    }
}

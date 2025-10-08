const std = @import("std");
const Lexer = @import("lexer/tokenize.zig");

pub fn main() !void {
    const allocator = std.heap.page_allocator;
    const path = "C:/Users/simon/AAAProjects/AAAFlowScript/test.flw";

    // otevře soubor
    const file = try std.fs.cwd().openFile(path, .{});
    defer file.close();

    // načte celý obsah
    const bytes = try file.readToEndAlloc(allocator, std.math.maxInt(usize));
    defer allocator.free(bytes);

    // výpis obsahu do konzole
    std.debug.print("{s}", .{bytes});
}

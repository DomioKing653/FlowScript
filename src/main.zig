const std = @import("std");
const Lexer = @import("lexer/tokenize.zig");
const Tokens = @import("lexer/tokens.zig");
const Parsing = @import("parser/parser.zig");

fn printer(text: []u8) !void {
    _ = try std.fs.File.stdout().write(text);
}
pub fn main() !void {
    const allocator = std.heap.page_allocator;
    const path = "C:/Users/simon/AAAProjects/AAAFlowScript/test.flw";

    // opening
    const file = try std.fs.cwd().openFile(try allocator.dupe(u8, path), .{});
    defer file.close();

    //reading file
    const bytes = try file.readToEndAlloc(allocator, std.math.maxInt(usize));
    defer allocator.free(bytes);

    var mainLexer = try Lexer.createLexer(bytes);

    const toks = mainLexer.tokenize() catch |err| {
        std.debug.print("Lexing error: {}\n", .{err});
        return;
    };
    for (toks) |tok| {
        std.debug.print("{s}->", .{try Tokens.TokenKindToString(tok.Kind)});
        switch (tok.Value) {
            .char => |c| std.debug.print("{c}\n", .{c}),
            .chars => |txt| std.debug.print("{s}\n", .{txt}),
        }
    }
    var mainParser = Parsing.Parser{ .tokens = try allocator.dupe(Tokens.Token, toks), .pos_idx = 0, .current_token = undefined, .statements = undefined, .alloc = allocator };
    mainParser.parse() catch |err| {
        try printer(try allocator.dupe(u8, @errorName(err)));
    };
}

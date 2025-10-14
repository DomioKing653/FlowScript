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
        std.debug.print("{s}\n", .{tok.Value});
    }
    var mainParser = Parsing.Parser{ .tokens = try allocator.dupe(Tokens.Token, toks), .pos_idx = 0, .current_token = undefined, .statements = undefined, .alloc = allocator };
    const stmts = try mainParser.parse();
    try printer(try allocator.dupe(u8, "AST:\n"));
    for (stmts) |stmt| {
        switch (stmt) {
            .varStmt => |v| std.debug.print("varStmt -> isConst: {any}, id: {s}, value: {?}\n", .{ v.is_const, v.id, v.value }),
        }
    }
}

const std = @import("std");
const tokens = @import("tokens.zig");

pub fn tokenize(code: []const u8) ![]tokens.Token {
    var charakterLoop: bool = false;
    const alloc = std.heap.page_allocator;
    var toks: std.ArrayList(tokens.Token) = .empty;
    defer toks.deinit(alloc);
    try std.fs.File.stdout().writeAll("[Lexing]...\n");
    for (code) |char| {
        switch (char) {
            '+' => {
                try toks.append(alloc, tokens.Token{ .Kind = tokens.TokenKind.PLUS, .Value = .{ .char = '+' } });
            },
            '-' => {
                try toks.append(alloc, tokens.Token{ .Kind = tokens.TokenKind.MINUS, .Value = .{ .char = '-' } });
            },
            '/' => {
                try toks.append(alloc, tokens.Token{ .Kind = tokens.TokenKind.SLASH, .Value = .{ .char = '/' } });
            },
            '*' => {
                try toks.append(alloc, tokens.Token{ .Kind = tokens.TokenKind.TIMES, .Value = .{ .char = '*' } });
            },
            else => {
                var id: std.ArrayList(u8) = .empty;
                if ((char >= 'a' and char <= 'z') or (char >= 'A' and char <= 'Z')) {
                    if (charakterLoop) {
                        try id.append(alloc, char);
                        charakterLoop = true;
                    } else {
                        charakterLoop = true;
                    }
                    continue;
                } else {
                    if (id.items.len == 0) {
                        const idStr = try id.toOwnedSlice(alloc);
                        try toks.append(alloc, tokens.Token{ .Kind = tokens.TokenKind.SYMBOL, .Value = .{ .chars = idStr } });
                        id.clearAndFree(alloc);
                    }
                    charakterLoop = false;
                }
            },
        }
    }
    return try toks.toOwnedSlice(alloc);
}

//imports
const std = @import("std");
const tokens = @import("tokens.zig");
//lexer variables
var posIdx: usize = 0;
var toks: std.ArrayList(tokens.Token) = .empty;
const alloc = std.heap.page_allocator;
var current_token: u8 = 89;
var code: []u8 = undefined;
//errors
const LexerError = error{UknowCharater};
//lexing
pub fn tokenize(program: []const u8) ![]tokens.Token {
    defer toks.deinit(alloc);
    code = try alloc.dupe(u8, program);

    current_token = code[0];
    posIdx += 1;
    try std.fs.File.stdout().writeAll("[Lexing]...\n");
    while (current_token != 0) {
        switch (current_token) {
            '+' => {
                try toks.append(alloc, tokens.Token{ .Kind = tokens.TokenKind.PLUS, .Value = .{ .char = '+' } });
                advance();
            },
            '-' => {
                try toks.append(alloc, tokens.Token{ .Kind = tokens.TokenKind.MINUS, .Value = .{ .char = '-' } });
                advance();
            },
            '/' => {
                try toks.append(alloc, tokens.Token{ .Kind = tokens.TokenKind.SLASH, .Value = .{ .char = '/' } });
                advance();
            },
            '*' => {
                try toks.append(alloc, tokens.Token{ .Kind = tokens.TokenKind.TIMES, .Value = .{ .char = '*' } });
                advance();
            },
            '=' => {
                try toks.append(alloc, tokens.Token{ .Kind = tokens.TokenKind.EQ, .Value = .{ .char = '=' } });
                advance();
            },
            '\n', ' ', '\t' => advance(),
            else => {
                if (std.ascii.isAlphabetic(current_token)) {
                    try toks.append(alloc, try lexSymbol());
                } else {
                    return LexerError.UknowCharater;
                }
            },
        }
    }
    //EOF
    const eof_arr: [3]u8 = [_]u8{ 'E', 'O', 'F' };
    const eof_text: []u8 = try alloc.dupe(u8, &eof_arr);
    try toks.append(alloc, tokens.Token{ .Kind = tokens.TokenKind.EOF, .Value = .{ .chars = eof_text } });
    //returns tokens
    return try toks.toOwnedSlice(alloc);
}

fn advance() void {
    if (posIdx >= code.len) {
        current_token = 0; // EOF token
        return;
    }
    current_token = code[posIdx];
    posIdx += 1;
}

fn lexSymbol() !tokens.Token {
    var text: std.ArrayList(u8) = .empty;
    while (std.ascii.isAlphabetic(current_token)) {
        try text.append(alloc, current_token);
        advance();
    }

    if (std.mem.eql(u8, text.items, "let")) {
        return tokens.Token{ .Kind = tokens.TokenKind.LET, .Value = .{ .chars = text.items } };
    } else {
        return tokens.Token{ .Kind = tokens.TokenKind.SYMBOL, .Value = .{ .chars = text.items } };
    }
}

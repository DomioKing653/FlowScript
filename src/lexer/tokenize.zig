//imports
const std = @import("std");
const tokens = @import("tokens.zig");

pub const Lexer: type = struct {
    pos_idx: usize,
    toks: std.ArrayList(tokens.Token),
    alloc: std.mem.Allocator,
    current_char: u8, //lexing
    code: []u8,
    pub fn tokenize(self: *Lexer) ![]tokens.Token {
        defer self.toks.deinit(self.alloc);
        self.current_char = self.code[0];
        self.pos_idx += 1;
        try std.fs.File.stdout().writeAll("[Lexing]...\n");
        while (self.current_char != 0) {
            switch (self.current_char) {
                '+' => {
                    try self.addToken(tokens.TokenKind.PLUS, .{ .char = '+' });
                },
                '-' => {
                    try self.addToken(tokens.TokenKind.MINUS, .{ .char = '-' });
                },
                '/' => {
                    self.advance();
                    if (self.current_char == '/') {
                        while (self.current_char != '\n') {
                            self.advance();
                        }
                        self.advance();
                    } else {
                        try self.addToken(tokens.TokenKind.SLASH, .{ .char = '+' });
                    }
                },
                '*' => {
                    try self.addToken(tokens.TokenKind.TIMES, .{ .char = '*' });
                    self.advance();
                },
                '=' => {
                    try self.addToken(tokens.TokenKind.EQ, .{ .char = '=' });
                    self.advance();
                },
                ';' => {
                    try self.addToken(tokens.TokenKind.SEMI_COLON, .{ .char = ';' });
                    self.advance();
                },
                '[' => {
                    try self.addToken(tokens.TokenKind.OPEN_BRACKET, .{ .char = ']' });
                    self.advance();
                },
                '\n', ' ', '\r', '\t' => self.advance(),
                else => {
                    if (std.ascii.isAlphabetic(self.current_char)) {
                        try self.lexSymbol();
                    } else {
                        return LexerError.UknowCharater;
                    }
                },
            }
        }
        //EOF
        try self.addToken(tokens.TokenKind.EOF, .{ .chars = try self.alloc.dupe(u8, "EOF") });
        //returns tokens
        return try self.toks.toOwnedSlice(self.alloc);
    }

    // Advancing
    fn advance(self: *Lexer) void {
        if (self.pos_idx >= self.code.len) {
            self.current_char = 0; // EOF token
            return;
        }
        self.current_char = self.code[self.pos_idx];
        self.pos_idx += 1;
    }

    //Symbol lexing
    fn lexSymbol(self: *Lexer) !void {
        var text: std.ArrayList(u8) = .empty;
        defer text.deinit(self.alloc);
        while (std.ascii.isAlphabetic(self.current_char) or self.current_char == '_') {
            try text.append(self.alloc, self.current_char);
            self.advance();
        }

        const dupe_text = try self.alloc.dupe(u8, text.items);
        if (std.mem.eql(u8, text.items, "let")) {
            try self.addToken(tokens.TokenKind.LET, .{ .chars = dupe_text });
        } else {
            try self.addToken(tokens.TokenKind.SYMBOL, .{ .chars = dupe_text });
        }
    }

    fn addToken(self: *Lexer, kind: tokens.TokenKind, value: tokens.value) !void {
        try self.toks.append(self.alloc, .{ .Kind = kind, .Value = value });
        self.advance();
    }
};

pub fn createLexer(code: []u8) !Lexer {
    const alloc = std.heap.page_allocator;
    return Lexer{ .alloc = alloc, .code = try alloc.dupe(u8, code), .current_char = undefined, .pos_idx = 0, .toks = undefined };
}

//errors
const LexerError = error{UknowCharater};
const UknowCharErr = struct { Err: LexerError, Char: u8 };

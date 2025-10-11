const std = @import("std");
pub const TokenKind: type = enum {
    //Math
    PLUS,
    MINUS,
    TIMES,
    SLASH,
    EQ,
    // Positional & Grouping
    SEMI_COLON,
    OPEN_PAREN,
    CLOSE_PAREN,
    OPEN_BRACKET,
    CLOSE_BRACKET,
    OPEN_CURLY,
    CLOSE_CURLY,
    //Values
    SYMBOL,
    NUMBER,
    //Keywords
    LET,
    CONST,
    //Parser
    EOF,
};

const allocator = std.heap.page_allocator;

pub const value = union(enum) { chars: []u8, char: u8 };

pub const Token: type = struct { Kind: TokenKind, Value: []u8 };

pub fn TokenKindToString(kind: TokenKind) ![]u8 {
    switch (kind) {
        //Math
        TokenKind.PLUS => return try allocator.dupe(u8, "PLUS"),
        TokenKind.MINUS => return try allocator.dupe(u8, "MINUS"),
        TokenKind.TIMES => return try allocator.dupe(u8, "TIMES"),
        TokenKind.SLASH => return try allocator.dupe(u8, "DIVIDE"),
        TokenKind.EQ => return try allocator.dupe(u8, "EQUAL"),
        //Keywords
        TokenKind.LET => return try allocator.dupe(u8, "LET"),
        TokenKind.CONST => return try allocator.dupe(u8, "CONST"),
        //Misc
        TokenKind.SYMBOL => return try allocator.dupe(u8, "IDENTIFIER"),
        TokenKind.SEMI_COLON => return try allocator.dupe(u8, "SEMI_COLON"),
        else => return try allocator.dupe(u8, "UNKNOWN"),
    }
}

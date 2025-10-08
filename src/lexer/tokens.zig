pub const TokenKind: type = enum {
    //Math
    PLUS,
    MINUS,
    TIMES,
    SLASH,
    //Values
    SYMBOL,
    NUMBER,
    //Parser
    EOF,
};

const value = union(enum) { chars: []u8, char: u8 };

pub const Token: type = struct { Kind: TokenKind, Value: value };

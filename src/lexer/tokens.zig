const tokenKind: type = enum {
    PLUS,
    MINUS,
    DASH,
    SLASH,
    SYMBOL,
    NUMBER,

    EOF,
};

const Token: type = struct { Kind: tokenKind, Value: []u8 };

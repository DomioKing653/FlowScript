const tokenKind: type = enum {
    //Math
    PLUS,
    MINUS,
    DASH,
    SLASH,
    //Values
    SYMBOL,
    NUMBER,
    //Parser
    EOF,
};

const Token: type = struct { Kind: tokenKind, Value: []u8 };

const Expression = @import(".//expression.zig");

pub const Statement = union(enum) { varStmt: vatDeclStmt };

const vatDeclStmt = struct { is_const: bool, id: []u8, value: ?Expression.Prefix_Expr };

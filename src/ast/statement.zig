pub const Statement = union { varStmt: vatDeclStmt };

const vatDeclStmt = struct { is_const: bool, id: []u8 };

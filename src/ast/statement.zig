const Expression = @import(".//expression.zig");
const Runtime = @import("../runtime/runtimeVal.zig");
pub const Statement = union(enum) {
    varStmt: varDeclStmt,
    fn stmt(self: *Statement) !Runtime.runtimeValue {
        switch (self) {
            .varStmt => return var_decl_visitnode(),
            else => unreachable,
        }
    }
};
//Statement interface
pub const IStatement = struct {
    ptr: *anyopaque,
    execFn: *const fn (ptr: *anyopaque) void,
    // Statement inicialization
    pub fn init(pointer: anytype) IStatement {
        const T = @TypeOf(pointer);

        const Gen = struct {
            fn opaqueExec(ptr: *anyopaque) void {
                const self: T = @ptrCast(@alignCast(ptr));
                self.exec();
            }
        };

        return IStatement{
            .ptr = pointer,
            .execFn = Gen.opaqueExec,
        };
    }

    pub fn exec(self: IStatement) void {
        self.execFn(self.ptr);
    }
};

pub const varDeclStmt = struct {
    is_const: bool,
    id: []u8,
    value: ?Expression.Prefix_Expr,
    pub fn exec(self: *varDeclStmt) void {
        _ = self;
    }
};

fn var_decl_visitnode() !Runtime.runtimeValue {}

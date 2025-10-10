const ast = @import("../ast/expression.zig");

test "ast_test" {
    const expresion = ast.BinaryExpr{.left};
    expresion.expr();
}

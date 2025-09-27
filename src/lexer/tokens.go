package lexer

type TokenKind int

const (
	EOF TokenKind = iota
	NUMBER
	STRING

	OPEN_BRACKET
	CLOSE_BRACKET
	OPEN_CURLY
	CLOSE_CURLY
	OPEN_PAREN
	CLOSE_PAREN

	PLUS
	MINUS
	DIVIDE
	MULTIPLY

	ASSGMENT // =
	EQUALS   // ==
	NOT
	NOT_EQUALS

	AND
	OR

	DOT
	DOT_DOT
)

type Token struct {
	Kind  TokenKind
	Value string
}

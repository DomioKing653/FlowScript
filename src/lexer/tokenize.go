package lexer

import (
	"fmt"
	"regexp"
)

type regexHandler func(lex *Lexer, regex *regexp.Regexp)

type RegexPattern struct {
	regex   *regexp.Regexp
	handler regexHandler
}

// Lexer holds the state for tokenizing source code
type Lexer struct {
	pattern    []RegexPattern
	Tokens     []Token
	source     string
	currentIdx int
}

// advanceN moves the current index forward by n positions
func (lex *Lexer) advanceN(n int) {
	lex.currentIdx += n
}

// push adds a token to the tokens list
func (lex *Lexer) push(tok Token) {
	lex.Tokens = append(lex.Tokens, tok)
}

// at_eof checks if we've reached the end of the source
func (lex *Lexer) at_eof() bool {
	return lex.currentIdx >= len(lex.source)
}

func (lex *Lexer) remainder() string {
	return lex.source[lex.currentIdx:]
}

// CreateLexer initializes a new lexer with the given source
func CreateLexer(source string) *Lexer {
	return &Lexer{
		source:     source,
		currentIdx: 0,
		Tokens:     make([]Token, 0),
		pattern: []RegexPattern{
			//numbers strings etc.
			{regexp.MustCompile(`[a-zA-Z_][a-zA-Z0-9_]*`), symbolHandler},
			{regexp.MustCompile(`[0-9]+(\.[0-9]+)?`), numberHandler},
			{regexp.MustCompile(`\s+`), skipHandler},
			{regexp.MustCompile(`\/\/.*`), skipHandler},
			{regexp.MustCompile(`"[^"]*"`), stringHandler},
			//default handlers
			{regexp.MustCompile(`\[`), defaultHandler(OPEN_BRACKET, "[")},
			{regexp.MustCompile(`\]`), defaultHandler(CLOSE_BRACKET, "]")},
			{regexp.MustCompile(`\{`), defaultHandler(OPEN_CURLY, "{")},
			{regexp.MustCompile(`\}`), defaultHandler(CLOSE_CURLY, "}")},
			{regexp.MustCompile(`\(`), defaultHandler(OPEN_PAREN, "(")},
			{regexp.MustCompile(`\)`), defaultHandler(CLOSE_PAREN, ")")},
			{regexp.MustCompile(`==`), defaultHandler(EQUALS, "==")},
			{regexp.MustCompile(`!=`), defaultHandler(NOT_EQUALS, "!=")},
			{regexp.MustCompile(`=`), defaultHandler(ASSIGNMENT, "=")},
			{regexp.MustCompile(`!`), defaultHandler(NOT, "!")},
			{regexp.MustCompile(`<=`), defaultHandler(LESS_EQUALS, "<=")},
			{regexp.MustCompile(`<`), defaultHandler(LESS, "<")},
			{regexp.MustCompile(`>=`), defaultHandler(GREATER_EQUALS, ">=")},
			{regexp.MustCompile(`>`), defaultHandler(GREATER, ">")},
			{regexp.MustCompile(`\|\|`), defaultHandler(OR, "||")},
			{regexp.MustCompile(`&&`), defaultHandler(AND, "&&")},
			{regexp.MustCompile(`\.\.`), defaultHandler(DOT_DOT, "..")},
			{regexp.MustCompile(`\.`), defaultHandler(DOT, ".")},
			{regexp.MustCompile(`;`), defaultHandler(SEMI_COLON, ";")},
			{regexp.MustCompile(`:`), defaultHandler(COLON, ":")},
			{regexp.MustCompile(`\?\?=`), defaultHandler(NULLISH_ASSIGNMENT, "??=")},
			{regexp.MustCompile(`\?`), defaultHandler(QUESTION, "?")},
			{regexp.MustCompile(`,`), defaultHandler(COMMA, ",")},
			{regexp.MustCompile(`\+\+`), defaultHandler(PLUS_PLUS, "++")},
			{regexp.MustCompile(`--`), defaultHandler(MINUS_MINUS, "--")},
			{regexp.MustCompile(`\+=`), defaultHandler(PLUS_EQUALS, "+=")},
			{regexp.MustCompile(`-=`), defaultHandler(MINUS_EQUALS, "-=")},
			{regexp.MustCompile(`\+`), defaultHandler(PLUS, "+")},
			{regexp.MustCompile(`-`), defaultHandler(DASH, "-")},
			{regexp.MustCompile(`/`), defaultHandler(SLASH, "/")},
			{regexp.MustCompile(`\*`), defaultHandler(STAR, "*")},
			{regexp.MustCompile(`%`), defaultHandler(PERCENT, "%")},
		},
	}
}

func Tokenize(source string) ([]Token, error) {
	lex := CreateLexer(source)
	for !lex.at_eof() {
		matched := false

		for _, pattern := range lex.pattern {
			loc := pattern.regex.FindStringIndex(lex.remainder())
			if loc != nil && loc[0] == 0 {
				pattern.handler(lex, pattern.regex)
				matched = true
				break
			}
		}
		if !matched {
			return []Token{}, fmt.Errorf("Error::Lexer->unknown token: %s", lex.remainder())
		}
	}
	lex.push(NewToken(EOF, "EOF"))
	return lex.Tokens, nil
}

/*
Handlers
*/

func defaultHandler(kind TokenKind, value string) regexHandler {
	return func(lex *Lexer, regex *regexp.Regexp) {
		lex.advanceN(len(value))
		lex.push(NewToken(kind, value))
	}
}

func numberHandler(lex *Lexer, regex *regexp.Regexp) {
	match := regex.FindString(lex.remainder())
	lex.push(NewToken(NUMBER, match))
	lex.advanceN(len(match))
}

func skipHandler(lex *Lexer, regex *regexp.Regexp) {
	match := regex.FindStringIndex(lex.remainder())
	lex.advanceN(match[1])
}

func stringHandler(lex *Lexer, regex *regexp.Regexp) {
	match := regex.FindStringIndex(lex.remainder())
	stringLiteral := lex.remainder()[match[0]+1 : match[1]-1]

	lex.push(NewToken(STRING, stringLiteral))
	lex.advanceN(len(stringLiteral) + 2)
}

func symbolHandler(lex *Lexer, regex *regexp.Regexp) {
	value := regex.FindString(lex.remainder())

	if kind, exists := keywords[value]; exists {
		lex.push(NewToken(kind, value))
	} else {
		lex.push(NewToken(IDENTIFIER, value))
	}

	lex.advanceN(len(value))
}

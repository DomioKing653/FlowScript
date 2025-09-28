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

func (lex *Lexer) at() byte {
	return lex.source[lex.currentIdx]
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
			{regexp.MustCompile(`[0-9]+(\.[0-9]+)?`), numberHandler},
			//default handlers
			{regexp.MustCompile(`\[`), defaultHandler(OPEN_BRACKET, "[")},
			{regexp.MustCompile(`\]`), defaultHandler(OPEN_BRACKET, "]")},
			{regexp.MustCompile(`\{`), defaultHandler(OPEN_CURLY, "{")},
			{regexp.MustCompile(`\}`), defaultHandler(CLOSE_CURLY, "}")},
			{regexp.MustCompile(`\(`), defaultHandler(OPEN_PAREN, "(")},
			{regexp.MustCompile(`\)`), defaultHandler(CLOSE_PAREN, ")")},
			{regexp.MustCompile(`\==`), defaultHandler(EQUALS, "==")},
			{regexp.MustCompile(`\!=`), defaultHandler(NOT_EQUALS, "!=")},
			{regexp.MustCompile(`\+=`), defaultHandler(PLUS_EQUALS, "+=")},
			{regexp.MustCompile(`\+`), defaultHandler(PLUS, "+")},
			{regexp.MustCompile(`-=`), defaultHandler(MINUS_EQUALS, "-=")},
			{regexp.MustCompile(`-`), defaultHandler(DASH, "-")},
		},
	}
}

func Tokenize(source string) []Token {
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
			panic(fmt.Sprintf("Error::Lexer->unknown token: %s", lex.remainder()))
		}
	}
	lex.push(NewToken(EOF, "EOF"))
	return lex.Tokens
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

package strict

import (
	"unicode"
)

type Token struct {
	Type  string
	Value string
}

type Lexer struct {
	source   string
	location int
}

func (lex *Lexer) Lex() ([]Token, err) {
	tokens := []Token{}

	for !lex.End() {
		c := lex.Peek()
		if c == '{' {
			tokens = append(tokens, Token{"SCOPE_START", "{"})
		} else if c == '}' {
			tokens = append(tokens, Token{"SCOPE_END", "}"})
		}

		lex.Next()
	}

	return tokens, nil
}

func (lex *Lexer) statement() ([]Token, bool) {
	start := lex.location
	ok, name := lex.variableName()
	if !ok {
		return false, name
	}

}

func (lex *Lexer) variableName() (Token, bool) {
	start := lex.location
	var name string

	for !lex.End() {
		c := lex.Peek()
		if unicode.IsDigit(c) || unicode.IsNumber(c) {
			name += c
			lex.Next()
			continue
		}

		return Token{}, false
	}

	return Token{"VARIABLE_NAME", name}, true
}

func (lex *Lexer) End() bool {
	if lex.location > len(lex.source) {
		return true
	}

	return false
}

func (lex *Lexer) Next() {
	lex.location += 1
}

func (lex *Lexer) Peek() rune {
	return rune(lex.source[lex.location])
}

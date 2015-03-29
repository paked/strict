package strict

import (
	"fmt"
	"unicode"
)

type TokenType string

const (
	ScopeStart   TokenType = "SCOPE_START"
	ScopeEnd     TokenType = "SCOPE_END"
	ListStart    TokenType = "LIST_START"
	ListEnd      TokenType = "LIST_END"
	Sender       TokenType = "SENDER"
	String       TokenType = "STRING"
	Number       TokenType = "NUMBER"
	VariableName TokenType = "VAR_NAME"
	Assign       TokenType = "ASSIGN"
)

type Token struct {
	Type  TokenType
	Value string
}

func (t Token) String() string {
	var postfix string
	if len(t.Value) > 1 {
		postfix = ":" + t.Value
	}
	return string(t.Type) + postfix
}

type Lexer struct {
	source   string
	location int
}

func (lex *Lexer) Lex() ([]Token, error) {
	fmt.Println(lex.source)
	tokens := []Token{}

	for !lex.End() {
		if !lex.skipSpace() {
			break
		}

		c := lex.Peek()

		if c == '{' {
			tokens = append(tokens, Token{ScopeStart, "{"})
		} else if c == '}' {
			tokens = append(tokens, Token{ScopeEnd, "}"})
		} else if toks, ok := lex.list(); ok {
			for _, token := range toks {
				tokens = append(tokens, token)
			}
		} else if toks, ok := lex.assignment(); ok {
			for _, token := range toks {
				tokens = append(tokens, token)
			}
		} else if token, ok := lex.sender(); ok {
			tokens = append(tokens, token)
		} else if token, ok := lex.name(); ok {
			tokens = append(tokens, token)
			continue
		}
		lex.Next()
	}

	return tokens, nil
}

func (lex *Lexer) skipSpace() bool {
	for !lex.End() {
		c := lex.Peek()
		if !unicode.IsSpace(c) {
			return true
		}

		lex.Next()
	}

	return false
}

func (lex *Lexer) sender() (Token, bool) {
	old := lex.location
	c := lex.Peek()
	if c != '>' {
		return Token{}, false
	}

	lex.Next()

	if c != '>' {
		lex.location = old
		return Token{}, false
	}

	return Token{Sender, ">>"}, true
}

func (lex *Lexer) space() {
	for !lex.End() {
		if c := lex.Peek(); !unicode.IsSpace(c) {
			break
		}
		fmt.Println("this isn't space:", string(lex.Peek()))
		lex.Next()
	}
}

func (lex *Lexer) assignment() ([]Token, bool) {
	tokens := []Token{}

	c := lex.Peek()
	if c != '=' {
		return tokens, false
	}

	lex.Next()
	tokens = append(tokens, Token{Assign, "="})

	lex.skipSpace()

	value, ok := lex.variable()
	if !ok {
		return []Token{}, false
	}

	tokens = append(tokens, value)

	return tokens, true
}

func (lex *Lexer) variable() (Token, bool) {
	if t, ok := lex.string(); ok {
		return t, ok
	} else if t, ok := lex.number(); ok {
		return t, ok
	} else if t, ok := lex.name(); ok {
		return t, ok
	}

	return Token{}, false
}

func (lex *Lexer) name() (Token, bool) {
	var content string
	for !lex.End() {
		c := lex.Peek()
		if !isAlphabet(c) {
			if content == "" {
				break
			}
			fmt.Println(lex.source[lex.location:])
			return Token{VariableName, content}, true
		}
		content += string(c)
		lex.Next()
	}

	return Token{}, false
}

func (lex *Lexer) number() (Token, bool) {
	var content string
	for !lex.End() {
		c := lex.Peek()
		if c < '0' || c > '9' {
			if content == "" {
				break
			}

			return Token{Number, content}, true
		}

		content += string(c)
		lex.Next()
	}

	return Token{}, false
}

func (lex *Lexer) list() ([]Token, bool) {
	old := lex.location

	tokens := []Token{}
	start, ok := lex.listStart()
	if !ok {
		fmt.Println("no list start", string(lex.Peek()))
		return tokens, false
	}

	tokens = append(tokens, start)

	for !lex.End() {
		value, ok := lex.variable()
		if !ok {
			return tokens, false
		}

		tokens = append(tokens, value)
		lex.space()
		c := lex.Peek()

		fmt.Println("char:", string(lex.Peek()))
		if c != ',' {
			end, ok := lex.listEnd()
			if !ok {
				lex.location = old
				return tokens, false
			}

			tokens = append(tokens, end)
			lex.Next()
			return tokens, true
		}
		lex.Next()
		lex.space()
	}

	return tokens, false
}

func (lex *Lexer) listStart() (Token, bool) {
	if lex.Peek() != '[' {
		return Token{}, false
	}

	fmt.Println("STARTING ARRAY")
	lex.Next()

	return Token{ListStart, "["}, true
}

func (lex *Lexer) string() (Token, bool) {
	if c := lex.Peek(); c != '"' {
		return Token{}, false
	}

	lex.Next()
	var content string
	for !lex.End() {
		c := lex.Peek()
		if c == '"' {
			lex.Next()
			fmt.Println(content)
			return Token{String, content}, true
		}

		content += string(c)
		lex.Next()
	}

	return Token{}, false
}

func (lex *Lexer) listEnd() (Token, bool) {
	if lex.Peek() != ']' {
		return Token{}, false
	}

	lex.Next()

	return Token{ListEnd, "]"}, true
}

func (lex *Lexer) End() bool {
	if lex.location >= len(lex.source) {
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

func isAlphabet(c rune) bool {
	if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
		return true
	}

	return false
}

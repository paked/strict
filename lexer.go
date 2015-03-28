package strict

import (
	"fmt"
	"unicode"
)

type Token struct {
	Type  string
	Value string
}

func (t Token) String() string {
	var postfix string
	if len(t.Value) > 1 {
		postfix = ":" + t.Value
	}
	return t.Type + postfix
}

type Lexer struct {
	source   string
	location int
}

func (lex *Lexer) Lex() ([]Token, error) {
	fmt.Println(lex.source)
	tokens := []Token{}

	for !lex.End() {
		c := lex.Peek()
		if c == '{' {
			tokens = append(tokens, Token{"SCOPE_START", "{"})
		} else if c == '}' {
			tokens = append(tokens, Token{"SCOPE_END", "}"})
		} else if toks, ok := lex.list(); ok {
			for _, token := range toks {
				tokens = append(tokens, token)
			}
		} else if token, ok := lex.sender(); ok {
			tokens = append(tokens, token)
		}
		lex.space()
		lex.Next()
	}

	return tokens, nil
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

	return Token{"SENDER", ">>"}, true
}

func (lex *Lexer) space() {
	for !lex.End() {
		if c := lex.Peek(); !unicode.IsSpace(c) {
			break
		}
		fmt.Println("char:", string(lex.Peek()))
		lex.Next()
	}
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
		value, ok := lex.string()
		if !ok {
			fmt.Println("no value")
			return tokens, false
			// no more list values -- something went wrong...
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

	return Token{"LIST_START", "["}, true
}

func (lex *Lexer) listValue() (Token, bool) {
	var content string
	for !lex.End() {
		c := lex.Peek()
		fmt.Println(string(c))
		if unicode.IsDigit(c) || unicode.IsLetter(c) {
			content += string(c)
			lex.Next()
			continue
		}

		ok := true
		if content == "" {
			fmt.Println("EMPTY CONTENT")
			ok = false
		}

		return Token{"LIST_VALUE", content}, ok
	}

	return Token{}, false
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
			return Token{"STRING", content}, true
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

	return Token{"LIST_END", "]"}, true
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

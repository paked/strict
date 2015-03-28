package strict

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestLexer(t *testing.T) {
	d, _ := ioutil.ReadFile("test.st")
	l := Lexer{}
	l.source = string(d)

	ts, err := l.Lex()
	fmt.Println(err)
	fmt.Println(ts)
}

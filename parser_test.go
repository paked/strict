package strict

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestVariableStores(t *testing.T) {
	fmt.Println("--TEST VARIABLE STORES--")
	root := NewVariableStore(nil)
	root.Set("a", Value{"set a in root"})

	if v, err := root.Get("a"); err == nil {
		fmt.Println("->>", v)
	}

	sub := NewVariableStore(&root)

	if v, err := sub.Get("a"); err == nil {
		fmt.Println("->> in sub:", v)
	}

	sub.Set("b", Value{"hehehe set in sub"})

	if v, err := sub.Get("b"); err == nil {
		fmt.Println("->> in sub:", v)
	}

	sub.Set("a", Value{"Hahaha set"})

	if v, err := sub.Get("a"); err == nil {
		if v.Contents != "Hahaha set" {
			t.Error("Didnt expect that content")
		}
	}
	if v, err := root.Get("a"); err == nil {
		if v.Contents != "Hahaha set" {
			t.Error("Didnt expect that content")
		}
		fmt.Println("got a:", v)
	}
}

func TestParse(t *testing.T) {
	fmt.Println("---parsing parising---")
	d, _ := ioutil.ReadFile("test.st")
	l := Lexer{}
	l.source = string(d)

	ts, _ := l.Lex()
	fmt.Println("--- PARSING 4 REALZIES ---")
	p := Parser{}
	p.source = ts

	r, err := p.parseLine()
	if err != nil {
		fmt.Println("SOMETHING FEFD UP!")
	}

	fmt.Println(r)
	store := NewVariableStore(nil)
	r.Eval(&store)
	fmt.Println(store)

	variable, err := store.Get("variable")
	fmt.Println(variable)
}

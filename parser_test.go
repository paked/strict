package strict

import (
	"fmt"
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

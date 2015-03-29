package strict

import (
	"errors"
	"fmt"
)

type Parser struct {
	source   []Token
	location int
}

func (p *Parser) parseScope() (Scope, error) {
	scope := NewScope(nil)
	t, err := p.Peek()
	if err != nil {
		return scope, err
	}

	if !t.IsType(ScopeStart) {
		return scope, err
	}
	p.Next()

	a, err := p.parseAssign()
	fmt.Println(err, a)
	scope.AddLine(a)

	p.Next()

	return scope, nil
}

func (p *Parser) parseAssign() (Runner, error) {
	a := Assignment{}

	name, err := p.Peek()
	if err != nil {
		fmt.Println("Naming exists")
		return nil, err
	}

	if !name.IsType(VariableName) {
		fmt.Println("didnt Found a variable name", name.Type)
		return nil, err
	}

	a.name = name.Value

	p.Next()

	sign, err := p.Peek()
	if err != nil {
		return nil, err
	}

	if !sign.IsType(Assign) {
		return nil, err
	}

	p.Next()

	value, err := p.Peek()
	if err != nil {
		return nil, err
	}

	if !value.IsType(String) {
		return nil, err
	}

	a.value = Value{value.Value}
	p.Next()

	return &a, nil
}

type Assignment struct {
	name  string
	value Value
}

func (a *Assignment) Eval(s *VariableStore) {
	s.Set(a.name, a.value)
}

func (p *Parser) parseLine() (Runner, error) {
	scope, err := p.parseScope()
	if err != nil {
		return nil, err
	}

	return &scope, nil
}

func (p *Parser) End() bool {
	if p.location >= len(p.source) {
		return true
	}

	return false
}

func (p *Parser) Peek() (Token, error) {
	if p.End() {
		return Token{}, errors.New("COuld not ge tthat token")
	}

	return p.source[p.location], nil
}

func (p *Parser) Next() {
	p.location += 1
}

type Runner interface {
	Eval(s *VariableStore)
}

func NewScope(parent *Scope) Scope {
	store := NewVariableStore(nil)
	if parent != nil {
		store.parent = parent.store
	}

	return Scope{parent: parent, store: &store, Contents: []Runner{}}
}

type Scope struct {
	Contents []Runner
	parent   *Scope
	store    *VariableStore
}

func (s *Scope) AddLine(r Runner) {
	s.Contents = append(s.Contents, r)
}

func (s *Scope) Eval(store *VariableStore) {
	for _, line := range s.Contents {
		line.Eval(store)
	}
}

func NewVariableStore(parent *VariableStore) VariableStore {
	return VariableStore{make(map[string]*Value), make(map[string]bool), parent}
}

type VariableStore struct {
	vars   map[string]*Value
	exists map[string]bool
	parent *VariableStore
}

func (vs *VariableStore) Get(key string) (*Value, error) {
	scope, err := vs.getScope(key)
	if err != nil {
		return nil, err
	}

	return scope.vars[key], nil
}

func (vs *VariableStore) Set(key string, value Value) {
	scope, err := vs.getScope(key)
	if err != nil {
		vs.exists[key] = true
		vs.vars[key] = &value
		return
	}

	scope.exists[key] = true
	scope.vars[key] = &value
	return
}

func (vs *VariableStore) getScope(key string) (*VariableStore, error) {
	if vs.exists[key] {
		return vs, nil
	}

	if vs.parent == nil {
		return nil, errors.New("That variable doesn't exist")
	}

	return vs.parent.getScope(key)

}

type Value struct {
	Contents string
}

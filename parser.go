package strict

import "errors"

type Parser struct {
	source   []Token
	location int
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
	Eval(s *Scope)
}

type Scope struct {
	Contents []Runner
	parent   *Scope
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

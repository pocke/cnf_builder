package cnf_builder

import "fmt"

// ----------- Var
type Var struct {
	id int
	b  bool
}

func (v *Var) Not() *Var {
	return &Var{
		id: v.id,
		b:  !v.b,
	}
}

// --------------- Builder
type Builder struct {
	vars    []*Var
	clauses [][]*Var
}

func New() *Builder {
	return &Builder{
		vars:    make([]*Var, 0),
		clauses: make([][]*Var, 0),
	}
}

func (b *Builder) NewVar() *Var {
	v := &Var{
		id: len(b.vars) + 1,
		b:  true,
	}

	b.vars = append(b.vars, v)
	return v
}

func (b *Builder) AddClause(c ...*Var) {
	b.clauses = append(b.clauses, c)
}

func (b *Builder) Build() string {
	res := fmt.Sprintf("p cnf %d %d\n", len(b.vars), len(b.clauses))
	for _, c := range b.clauses {
		for _, v := range c {
			if !v.b {
				res += "-"
			}
			res += fmt.Sprintf("%d ", v.id)
		}
		res += "0\n"
	}
	return res
}

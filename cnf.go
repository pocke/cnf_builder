package cnf_builder

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

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

func Import(cnf io.Reader) (*Builder, error) {
	sc := bufio.NewScanner(cnf)
	for {
		ok := sc.Scan()
		if !ok {
			return nil, fmt.Errorf("Unexpected EOF")
		}
		line := sc.Text()
		if !strings.HasPrefix(line, "c") {
			break
		}
	}
	pcnf := sc.Text()

	pcnfs := strings.Split(pcnf, " ")
	if pcnfs[0] != "p" || pcnfs[1] != "cnf" {
		return nil, fmt.Errorf("Expected \"p cnf nbvar nbclauses\", but got %s", pcnf)
	}
	nbvar, err := strconv.Atoi(pcnfs[2])
	if err != nil {
		return nil, err
	}
	nbclauses, err := strconv.Atoi(pcnfs[3])
	if err != nil {
		return nil, err
	}

	b := &Builder{
		vars:    make([]*Var, 0, nbvar),
		clauses: make([][]*Var, 0, nbclauses),
	}
	for i := 0; i < nbvar; i++ {
		b.NewVar()
	}

	for i := 0; i < nbclauses; i++ {
		ok := sc.Scan()
		if !ok {
			return nil, fmt.Errorf("Unexpected EOF")
		}
		line := strings.Split(sc.Text(), " ")
		if last := line[len(line)-1]; last != "0" {
			return nil, fmt.Errorf("Last of clauses line should be 0, but got %s", last)
		}
		cs := make([]*Var, 0, len(line)-1)
		for _, s := range line[:len(line)-1] {
			n, err := strconv.Atoi(s)
			if err != nil {
				return nil, err
			}
			if abs(n) > nbvar {
				return nil, fmt.Errorf("Invalid as CNF")
			}
			v := b.vars[abs(n)-1]
			if n < 0 {
				v = v.Not()
			}
			cs = append(cs, v)
		}
		b.AddClause(cs...)
	}

	return b, nil
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
	res := make([]byte, 0, 15+16*len(b.clauses))
	res = append(res, fmt.Sprintf("p cnf %d %d\n", len(b.vars), len(b.clauses))...)
	for _, c := range b.clauses {
		for _, v := range c {
			if !v.b {
				res = append(res, '-')
			}
			res = append(res, strconv.Itoa(v.id)...)
			res = append(res, ' ')
		}
		res = append(res, "0\n"...)
	}
	return string(res)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

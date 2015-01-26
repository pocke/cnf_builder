package main

import (
	"fmt"

	"github.com/pocke/cnf_builder"
)

func main() {
	b := cnf_builder.New()
	p1 := b.NewVar()
	p2 := b.NewVar()
	p3 := b.NewVar()

	np1 := p1.Not()
	np2 := p2.Not()
	np3 := p3.Not()

	b.AddClause(p1, p2, p3)
	b.AddClause(np1, np2)
	b.AddClause(np1, np3)
	b.AddClause(np2, np3)

	s := b.Build()
	fmt.Print(s)
}

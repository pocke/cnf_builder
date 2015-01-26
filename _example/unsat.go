package main

import (
	"fmt"

	"github.com/pocke/cnf_builder"
)

func main() {
	b := cnf_builder.New()
	p1 := b.NewVar()

	b.AddClause(p1)
	b.AddClause(p1.Not())

	s := b.Build()
	fmt.Print(s)
}

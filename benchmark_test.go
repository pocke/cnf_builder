package cnf_builder_test

import (
	"testing"

	"github.com/pocke/cnf_builder"
)

func BenchmarkBuild(b *testing.B) {
	for i := 0; i < b.N; i++ {
		builder.Build()
	}
}

var builder = func() *cnf_builder.Builder {
	b := cnf_builder.New()
	v1 := b.NewVar()
	v2 := b.NewVar().Not()
	v3 := b.NewVar()
	for i := 0; i < 100; i++ {
		b.AddClause(v1, v2, v3)
	}
	return b
}()

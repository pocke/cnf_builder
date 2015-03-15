package cnf_builder_test

import (
	"strings"
	"testing"

	"github.com/pocke/cnf_builder"
)

func TestAbs(t *testing.T) {
	i := cnf_builder.Abs(100)
	if i != 100 {
		t.Errorf("Expected 100, but got %d", i)
	}
	i = cnf_builder.Abs(-100)
	if i != 100 {
		t.Errorf("Expected 100, but got %d", i)
	}
}

func shouldEqBuild(b *cnf_builder.Builder, str string, t *testing.T) {
	builded := b.Build()
	if builded != str {
		t.Fatalf(`Expected:
%s
but got
%s`, str, builded)
	}
}

func TestBuild(t *testing.T) {
	b := cnf_builder.New()
	v := b.NewVar()
	b.AddClause(v)

	shouldEqBuild(b, `p cnf 1 1
1 0
`, t)

	v2 := b.NewVar()
	b.AddClause(v, v2.Not())
	shouldEqBuild(b, `p cnf 2 2
1 0
1 -2 0
`, t)
}

func TestImportBuild(t *testing.T) {
	cnf := `p cnf 3 2
1 -2 3 0
1 -3 0
`
	b, _ := cnf_builder.Import(strings.NewReader(cnf))
	shouldEqBuild(b, cnf, t)
}

func TestImport(t *testing.T) {
	shouldBeError := func(s string) {
		r := strings.NewReader(s)
		_, err := cnf_builder.Import(r)
		if err == nil {
			t.Fatal("Expected error, but got nil")
		}
	}

	shouldBeSuccess := func(s string) {
		r := strings.NewReader(s)
		_, err := cnf_builder.Import(r)
		if err != nil {
			t.Fatal(err)
		}
	}

	shouldBeError(`hogefuga`)
	shouldBeError(` `)
	shouldBeError(``)
	shouldBeError(`p hoge 2 2
1 2 0
-1 -2 0`)
	shouldBeError(`p cnf 2 2
1 c 0
-1 -2 0`)
	shouldBeError(`p cnf a 2
1 2 0
-1 -2 0`)
	shouldBeError(`p cnf 2 b
1 2 0
-1 -2 0`)
	shouldBeSuccess(`p cnf 2 2
1 2 0
-1 -2 0`)
	shouldBeError(`p cnf 1 1
1`)
	shouldBeSuccess(`p cnf 1 1
1 0`)
	shouldBeSuccess(`p cnf 2 3
1 -2 0
1 0
2 0`)
	shouldBeSuccess(`c nyaafdshfkjsahjfkjsakf
c 1111111
p cnf 3 2
1 -2 3 0
1 -3 0`)
	shouldBeError(`p cnf 2 3
1 -2 0
-1 2 0`)
	shouldBeSuccess(`p cnf 2 3
1 -2 0
-1 2 0
1 2 0`)
	shouldBeError(`p cnf 1 1
1 2 0`)
	shouldBeSuccess(`p cnf 1 1
1 -1 0`)
}

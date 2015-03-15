package cnf_builder_test

import (
	"strings"
	"testing"

	"github.com/pocke/cnf_builder"
)

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

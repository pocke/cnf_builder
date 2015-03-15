package main

import (
	"fmt"
	"math"

	"github.com/pocke/cnf_builder"
)

var N = 9

func main() {
	b := cnf_builder.New()

	vars := make([][][]*cnf_builder.Var, 0, N)
	for i := 0; i < N; i++ {
		vars = append(vars, make([][]*cnf_builder.Var, 0, N))
		for j := 0; j < N; j++ {
			vars[i] = append(vars[i], make([]*cnf_builder.Var, 0, N))
			for k := 0; k < N; k++ {
				vars[i][j] = append(vars[i][j], b.NewVar())
			}
		}
	}

	for _, line := range vars {
		for _, cell := range line {
			// 各マスは1..Nのいずれかの数字が入る
			b.AddClause(cell...)
		}
	}

	comb := Combination(N)

	// 各列に対して同じ数字が2回以上現れない
	for _, line := range vars {
		for num := 0; num < N; num++ {
			for _, c := range comb {
				b.AddClause(line[c[0]][num].Not(), line[c[1]][num].Not())
			}
		}
	}

	// 各行に対して同じ数字が2回以上現れない
	for i := 0; i < N; i++ {
		for num := 0; num < N; num++ {
			for _, c := range comb {
				b.AddClause(vars[c[0]][i][num].Not(), vars[c[1]][i][num].Not())
			}
		}
	}

	// 各ブロックに対して同じ数字が2回以上現れない
	m := int(math.Sqrt(float64(N)))
	for i := 0; i < m; i++ {
		for j := 0; j < m; j++ {
			for _, c := range comb {
				x1 := c[0] % m
				x2 := c[1] % m
				y1 := c[0] / m
				y2 := c[1] / m

				for num := 0; num < N; num++ {
					b.AddClause(vars[i*m+y1][j*m+x1][num].Not(), vars[i*m+y2][j*m+x2][num].Not())
				}
			}
		}
	}

	s := b.Build()
	fmt.Printf(s)
}

func Combination(n int) [][]int {
	var res = make([][]int, 0)

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			res = append(res, []int{i, j})
		}
	}

	return res
}

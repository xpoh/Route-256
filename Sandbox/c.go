package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
)

type Arr struct {
	A   *[][]int
	Col int
}

func (a *Arr) Len() int {
	return len(*a.A)
}

func (a *Arr) Less(i, j int) bool {
	return (*a.A)[i][a.Col] < (*a.A)[j][a.Col]
}

func (a *Arr) Swap(i, j int) {
	(*a.A)[i], (*a.A)[j] = (*a.A)[j], (*a.A)[i]
}

func solve(a *[][]int, c []int) {
	arr := Arr{
		A: a,
	}

	for _, v := range c {
		arr.Col = v
		sort.Sort(&arr)
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	var n, m int
	var k int

	for _, _ = Fscan(in, &t); t > 0; t-- {
		_, _ = Fscan(in, &n, &m)
		a := make([][]int, n)
		var p1 int
		for i := 0; i < n; i++ {
			a[i] = make([]int, m)
			for j := 0; j < m; j++ {
				_, _ = Fscan(in, &p1)
				a[i][j] = p1
			}
		}
		_, _ = Fscan(in, &k)
		c := make([]int, k)
		for i := 0; i < k; i++ {
			_, _ = Fscan(in, &p1)
			c[i] = p1 - 1
		}

		solve(&a, c)
		outResult(out, &a)
		_, _ = Fprintln(out, "")
	}
}

func outResult(out *bufio.Writer, a *[][]int) {
	n := len(*a)
	m := len((*a)[0])

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			_, _ = Fprint(out, (*a)[i][j])
			if j < m-1 {
				_, _ = Fprint(out, " ")
			}
		}
		_, _ = Fprint(out, "\n")
	}
}

package main

import (
	"bufio"
	. "fmt"
	"os"
)

func solve(p []int) int {
	m := make(map[int]int)
	for _, v := range p {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}
	s := 0
	for k, v := range m {
		n := v / 3
		s = s + k*(v-n)
	}
	return s
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	var n int
	var p1 int

	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n)
		p := make([]int, n)
		for i := 0; i < n; i++ {
			Fscan(in, &p1)
			p[i] = p1
		}
		Fprintln(out, solve(p))
	}
}

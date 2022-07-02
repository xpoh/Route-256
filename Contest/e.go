package main

import (
	"bufio"
	. "fmt"
	"os"
)

func SolveE(n int, a []int) string {
	last_job := -1
	m := make(map[int]struct{})
	for _, job := range a {
		if job != last_job {
			last_job = job
			if _, ok := m[job]; ok {
				return "NO"
			} else {
				m[job] = struct{}{}
			}
		}
	}

	return "YES"
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	var n int
	var a []int

	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n)
		a = make([]int, n)
		for i := 0; i < n; i++ {
			Fscan(in, &a[i])
		}
		Println(SolveE(n, a))
	}
}

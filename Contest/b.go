package main

import (
	"bufio"
	. "fmt"
	"math"
	"os"
)

func solve(n int, a []int) {
	var k int
	m := make(map[int]struct{})
	for i := 0; i < n-1; i++ {
		if _, ok := m[i]; ok {
			continue
		}
		t := 101
		for j := i + 1; j < n; j++ {
			if _, ok := m[j]; ok {
				continue
			}
			t1 := int(math.Abs(float64(a[i] - a[j])))
			if t1 < t {
				t = t1
				k = j
			}
		}
		m[k] = struct{}{}
		m[i] = struct{}{}

		Println(i+1, k+1)
	}
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
		solve(n, a)
	}
}

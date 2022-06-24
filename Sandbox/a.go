package main

import (
	"bufio"
	. "fmt"
	"os"
)

func solve(s int64, n int64) int64 {
	return s / (n * n)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var T int
	var s, n int64
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		Fprintln(out, solve(s, n))
	}
}

package main

import (
	"bufio"
	. "fmt"
	"os"
)

func solve(a int64, b int64) int64 {
	return a + b
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	var a, b int64
	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &a, &b)
		Fprintln(out, solve(a, b))
	}
}

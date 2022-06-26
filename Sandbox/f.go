package main

import (
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

func solve(v []bool, cmd, ind int) (bool, int) {
	ind = ind - 1
	if cmd == 1 {
		if !v[ind] {
			v[ind] = true
			return true, -1
		} else {
			return false, -1
		}
	}
	if cmd == 2 {
		if v[ind] {
			v[ind] = false
			return true, -1
		} else {
			return false, -1
		}
	}
	for i := 0; i < len(v); i = i + 2 {
		if (!v[i]) && (v[i] == v[i+1]) {
			v[i] = true
			v[i+1] = true
			return true, i + 1
		}
	}

	return false, -1
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	var n int
	var m int
	//var k int

	for _, _ = Fscan(in, &t); t > 0; t-- {
		_, _ = Fscan(in, &n, &m)
		v := make([]bool, 2*n)
		cmd := 0
		ind := 0
		b := false
		j := 0
		for i := 0; i < m; i++ {
			_, _ = Fscan(in, &cmd)
			if cmd < 3 {
				_, _ = Fscan(in, &ind)
			}
			b, j = solve(v, cmd, ind)
			if b {
				_, _ = Fprint(out, "SUCCESS")
			} else {
				_, _ = Fprint(out, "FAIL")
			}
			if j > 0 {
				_, _ = Fprint(out, " ", strconv.Itoa(j), "-", strconv.Itoa(j+1)+"\n")
			} else {
				_, _ = Fprint(out, "\n")
			}
		}
		_, _ = Fprintln(out, "")
	}
}

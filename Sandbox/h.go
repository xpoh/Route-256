package main

import (
	"bufio"
	. "fmt"
	"math"
	"os"
	"sort"
)

func NearOrthCount(A [][]string, i, j int) int {
	n := 0
	if A[i-1][j] == "*" {
		n++
	}
	if A[i+1][j] == "*" {
		n++
	}
	if A[i][j-1] == "*" {
		n++
	}
	if A[i][j+1] == "*" {
		n++
	}
	return n
}

func NearDiagCount(A [][]string, i, j int) int {
	n := 0
	if A[i-1][j-1] == "*" {
		n++
	}
	if A[i+1][j+1] == "*" {
		n++
	}
	if A[i+1][j-1] == "*" {
		n++
	}
	if A[i-1][j+1] == "*" {
		n++
	}
	return n
}

func getDirection(A [][]string, i, j int) (int, int, int, bool) {
	c, di, dj := 0, 0, 0
	if A[i-1][j] == "*" {
		di = -1
		c = 0
	}
	if A[i+1][j] == "*" {
		di = 1
		c = 1
	}
	if A[i][j-1] == "*" {
		dj = -1
		c = 2
	}
	if A[i][j+1] == "*" {
		dj = 1
		c = 3
	}
	return c, di, dj, (di == 0) && (dj == 0)
}

func solve(A [][]string) (string, []int) {
	var sz []int

	for i := 1; i < len(A)-1; i++ {
		for j := 1; j < len(A[i])-1; j++ {
			if (A[i][j] != ".") && (A[i][j] != "*") {
				return "NO", nil
			}
			if A[i][j] == "*" {
				n := NearOrthCount(A, i, j)
				N := NearDiagCount(A, i, j)
				ok := ((n == 1) && (N == 0)) || ((n == 2) && (N == 0)) || ((n == 1) && (N == 1)) || ((n == 0) && (N == 0)) || (n == 2) && (N == 1)
				if !ok {
					return "NO", nil
				}
				if (n == 2) && (N == 1) {
					ok = (A[i+1][j] == A[i-1][j]) || (A[i][j+1] == A[i][j-1])
					if !ok {
						return "NO", nil
					}
				}
				if (n == 1) && (N == 1) {
					if A[i+1][j] == "*" {
						ok = (A[i+1][j] == A[i+1][j+1]) || (A[i+1][j] == A[i+1][j-1])
					}
					if A[i-1][j] == "*" {
						ok = (A[i-1][j] == A[i-1][j+1]) || (A[i-1][j] == A[i-1][j-1])
					}
					if A[i][j+1] == "*" {
						ok = (A[i][j+1] == A[i-1][j+1]) || (A[i][j+1] == A[i+1][j+1])
					}
					if A[i][j-1] == "*" {
						ok = (A[i][j-1] == A[i-1][j-1]) || (A[i][j-1] == A[i+1][j-1])
					}
					if !ok {
						return "NO", nil
					}
				}
			}
		}
	}
	for i := 1; i < len(A)-1; i++ {
		for j := 1; j < len(A[i])-1; j++ {
			if A[i][j] == "*" {
				n := NearOrthCount(A, i, j)

				if n == 0 {
					sz = append(sz, 1)
				}

				if n == 1 {
					l := 1
					rotCount := 0
					i1, j1 := i, j
					dir, di, dj, stop := getDirection(A, i1, j1)
					if (di != 0) && (dj != 0) {
						return "NO", nil
					}
					A[i1][j1] = "."
					dirOld := dir
					for !stop {
						l++
						i1 = i1 + di
						j1 = j1 + dj

						dir, di, dj, stop = getDirection(A, i1, j1)
						if (di != 0) && (dj != 0) {
							return "NO", nil
						}

						A[i1][j1] = "."
						if (dir != dirOld) && (!stop) {
							rotCount++
						}
						dirOld = dir
					}
					if rotCount != 1 {
						return "NO", nil
					}
					if math.Abs(float64(i1-i)) != math.Abs(float64(j1-j)) {
						return "NO", nil
					}
					sz = append(sz, l)
				}
			}
		}
	}
	if sz == nil {
		return "NO", nil
	}
	sort.Ints(sz)
	return "YES", sz
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	var n int
	var m int

	for _, _ = Fscan(in, &t); t > 0; t-- {

		//Fprintf(out, "Test %v\n", t)
		//out.Flush()
		//if t == 75 {
		//	Fprintf(out, "DANGER!!!")
		//	out.Flush()
		//}

		_, _ = Fscan(in, &n)
		_, _ = Fscan(in, &m)

		A := make([][]string, n+2)
		var s string
		for i, _ := range A {
			A[i] = make([]string, m+2)
			A[i][0] = "."
			A[i][m+1] = "."
		}
		for j, _ := range A[0] {
			A[0][j] = "."
			A[n+1][j] = "."
		}

		for j := 0; j < n; j++ {
			_, _ = Fscan(in, &s)
			for i, c := range s {
				A[j+1][i+1] = string(c)
			}
		}

		ans, k := solve(A)

		//Fprintf(out, "Test %v - ", t)

		_, _ = Fprint(out, ans)

		if ans == "YES" {
			_, _ = Fprint(out, "\n")
			for i, v := range k {
				if i > 0 {
					_, _ = Fprint(out, " ")
				}
				_, _ = Fprint(out, v)
			}
			if t > 1 {
				_, _ = Fprint(out, "\n")
			}
		} else {
			if t > 1 {
				_, _ = Fprint(out, "\n")
			}
		}
	}
}

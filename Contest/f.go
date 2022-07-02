package main

import (
	"bufio"
	. "fmt"
	"io/fs"
	"os"
	"sort"
	"strconv"
	"strings"
)

func DateValid(h, m, s int) bool {
	return (h >= 0) && (h < 24) && (m >= 0) && (m <= 59) && (s >= 0) && (s <= 59)
}

func ParseString(s string) (a, b int, ok bool) {
	s1 := strings.Split(s, "-")
	d1 := strings.Split(s1[0], ":")
	d2 := strings.Split(s1[1], ":")
	var h, m, ss int
	var h1, m1, ss1 int
	h, _ = strconv.Atoi(d1[0])
	m, _ = strconv.Atoi(d1[1])
	ss, _ = strconv.Atoi(d1[2])
	h1, _ = strconv.Atoi(d2[0])
	m1, _ = strconv.Atoi(d2[1])
	ss1, _ = strconv.Atoi(d2[2])

	//if s == "23:59:59-23:59:58" {
	//	Println("DEBUG")
	//}

	if DateValid(h, m, ss) && DateValid(h1, m1, ss1) {
		a := 3600*h + 60*m + ss
		b := 3600*h1 + 60*m1 + ss1
		if a <= b {
			return a, b, true
		}
	}
	return 0, 0, false
}

type D struct {
	a, b []int
}

func (d *D) Len() int {
	return len(d.a)
}

func (d *D) Less(i, j int) bool {
	return d.a[i] < d.a[j]
}

func (d *D) Swap(i, j int) {
	d.a[i], d.a[j] = d.a[j], d.a[i]
	d.b[i], d.b[j] = d.b[j], d.b[i]
}

func Cross(a, b []int) string {
	//for i := 0; i < len(a)-1; i++ {
	//	for j := i + 1; j < len(a); j++ {
	//		if ((a[j] <= b[i]) && (a[j] >= a[i])) || ((b[j] <= b[i]) && (b[j] >= a[i])) {
	//			return "NO"
	//		}
	//	}
	//}
	//return "YES"
	d := D{
		a: a,
		b: b,
	}
	sort.Sort(&d)
	for i := 0; i < len(a)-1; i++ {
		if d.b[i] >= d.a[i+1] {
			return "NO"
		}
	}
	return "YES"
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	f, _ := os.OpenFile("f.out", os.O_CREATE, fs.ModePerm)
	defer out.Flush()
	defer f.Close()

	var t int
	var n int
	var s string
	var ok bool = true

	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &n)

		a := make([]int, n)
		b := make([]int, n)
		ok = true
		for i := 0; i < n; i++ {
			Fscan(in, &s)
			if ok {
				a[i], b[i], ok = ParseString(s)
				if !ok {
					Println("NO")
					Fprintln(f, "NO")
				}
			}
		}
		if ok && (len(a) == 1) {
			Println("YES")
			Fprintln(f, "YES")
		}
		if ok && (len(a) > 1) {
			ans := Cross(a, b)
			Println(ans)
			Fprintln(f, ans)
		}
	}
}

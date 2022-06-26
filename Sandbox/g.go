package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

var cache map[string]struct{}
var m map[string][]string

func solve(module string) (int, string) {
	if _, ok := cache[module]; ok {
		return 0, ""
	}

	ans := module
	n := 1
	cache[module] = struct{}{}

	if (len(m[module]) == 1) && (m[module][0] == "") {
		return n, ans
	}

	for _, s := range m[module] {
		if s == "" {
			continue
		}
		if _, ok := cache[s]; ok {
			continue
		} else {
			cache[s] = struct{}{}
			ans = s + " " + ans
		}
	}

	for _, s := range m[module] {
		if s == "" {
			continue
		}

		n1, ans1 := solve(s)
		if n1 > 0 {
			n = n + n1
			ans = ans1 + " " + ans
		}
	}
	return n, ans
}

func addModule(s string) {
	s = strings.TrimRight(s, "\r\n")
	s1 := strings.Split(s, ":")
	s2 := strings.TrimLeft(s1[1], " ")
	s3 := strings.Split(s2, " ")
	m[s1[0]] = s3
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	var n int
	var q int
	var s string

	for _, _ = Fscan(in, &t); t > 0; t-- {
		m = make(map[string][]string)
		cache = make(map[string]struct{})

		_, _ = Fscanln(in)
		_, _ = Fscan(in, &n)
		s, _ = in.ReadString('\n')
		for i := 0; i < n; i++ {
			s, _ = in.ReadString('\n')
			addModule(s)
		}
		_, _ = Fscan(in, &q)
		for i := 0; i < q; i++ {
			_, _ = Fscan(in, &s)
			n, s = solve(s)
			_, _ = Fprintf(out, "%d", n)
			if n > 0 {
				_, _ = Fprintf(out, " %s", s)
			}
			//if i < q-1 {
			_, _ = Fprint(out, "\n")
			//}
		}
		if t > 1 {
			_, _ = Fprintln(out, "")
		}
	}
}

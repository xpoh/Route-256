package main

import (
	"bufio"
	. "fmt"
	"os"
	"sort"
	"strconv"
)

func uniq(s []string) []string {
	m := make(map[string]struct{})
	var ans []string
	for _, v := range s {
		m[v] = struct{}{}
	}
	for k, _ := range m {
		ans = append(ans, k)
	}
	return ans
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	var n int
	var s1 string
	var s2 string

	for _, _ = Fscan(in, &t); t > 0; t-- {
		_, _ = Fscan(in, &n)
		book := make(map[string][]string)
		names := make([]string, 0)

		for i := 0; i < n; i++ {
			_, _ = Fscan(in, &s1)
			_, _ = Fscan(in, &s2)
			book[s1] = append(book[s1], s2)
		}
		for k, _ := range book {
			names = append(names, k)
		}
		sort.Strings(names)
		for _, name := range names {
			history := make(map[string]struct{})

			sl := book[name]
			l := len(sl)
			ans := make([]string, 0)

			for i, _ := range sl {
				if _, ok := history[sl[l-i-1]]; !ok {
					ans = append(ans, sl[l-i-1])
					history[sl[l-i-1]] = struct{}{}
				}
			}
			l = len(ans)
			if l > 5 {
				l = 5
			}
			_, _ = Fprint(out, name, ":", " ", strconv.Itoa(l))
			for i := 0; i < l; i++ {
				_, _ = Fprint(out, " ", ans[i])
			}
			_, _ = Fprint(out, "\n")
		}
		if t > 1 {
			_, _ = Fprint(out, "\n")
		}
	}
}

package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

var mask string = "qwertyuiopasdfghjklzxcvbnm-_1234567890"
var m map[string]struct{}

func solve(s string) bool {
	// test prefix
	if strings.HasPrefix(s, "-") {
		return false
	}
	// test lenght
	l := len(s)

	if (l < 2) || (l > 24) {
		return false
	}

	// TEST MASK
	slo := strings.ToLower(s)
	for _, c := range slo {
		if !strings.Contains(mask, string(c)) {
			return false
		}
	}
	// test first
	if _, ok := m[slo]; ok {
		return false
	} else {
		m[slo] = struct{}{}
	}

	return true
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	//f, _ := os.OpenFile("d.out", os.O_CREATE, fs.ModePerm)
	defer out.Flush()

	var t int32
	var n int32
	//var m int
	//var k int

	for _, _ = Fscan(in, &t); t > 0; t-- {
		_, _ = Fscan(in, &n)
		m = make(map[string]struct{})
		var s string
		for i := int32(0); i < n; i++ {
			_, _ = Fscan(in, &s)
			if solve(s) {
				_, _ = Fprintln(out, "YES")
				//_, _ = Fprintln(f, "YES"+" "+s)

			} else {
				//_, _ = Fprintln(f, "NO"+" "+s)
				_, _ = Fprintln(out, "NO")
			}
		}
		_, _ = Fprintln(out, "")
	}
}

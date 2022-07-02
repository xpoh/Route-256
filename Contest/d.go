package main

import (
	"bufio"
	. "fmt"
	"os"
	"strings"
)

var s_glas string = "euioay"
var s_soglas string = "qwrtplkjhgfdszxcvbnm"
var s_num string = "1234567890"

func solve(s string) string {
	ans := s
	s_lo := strings.ToLower(s)
	s_hi := strings.ToUpper(s)
	var t1, t2, t3, t4, t5 bool
	//1 хотя бы одну гласную букву;
	for _, c := range s_glas {
		for j, c1 := range s_lo {
			if c == c1 {
				t1 = true
				//3 хотя бы одну прописную букву (букву в верхнем регистре);
				if s[j] == s_hi[j] {
					t3 = true
				}
				//4 хотя бы одну строчную букву (букву в нижнем регистре);
				if s[j] == s_lo[j] {
					t4 = true
				}
			}
		}
	}
	//2 хотя бы одну согласную букву;
	for _, c := range s_soglas {
		for j, c1 := range s_lo {
			if c == c1 {
				t2 = true
				//3 хотя бы одну прописную букву (букву в верхнем регистре);
				if s[j] == s_hi[j] {
					t3 = true
				}
				//4 хотя бы одну строчную букву (букву в нижнем регистре);
				if s[j] == s_lo[j] {
					t4 = true
				}
			}
		}
	}

	//5 хотя бы одну цифру.
	for _, c := range s_num {
		for _, c1 := range s_lo {
			if c == c1 {
				t5 = true
				break
			}
		}
		if t5 {
			break
		}
	}
	if t5 == false {
		ans = ans + "7"
	}

	if t1 == false {
		if t3 == false {
			ans = ans + "A"
			t3 = true
		} else if t4 == false {
			ans = ans + "a"
			t4 = true
		} else {
			ans = ans + "e"
		}
	}

	if t2 == false {
		if t3 == false {
			ans = ans + "R"
			t3 = true
		} else if t4 == false {
			ans = ans + "r"
			t4 = true
		} else {
			ans = ans + "t"
		}
	}
	if t3 == false {
		ans = ans + "T"
	}
	if t4 == false {
		ans = ans + "t"
	}

	return ans
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	var s string

	for Fscan(in, &t); t > 0; t-- {
		Fscan(in, &s)
		Println(solve(s))
	}
}

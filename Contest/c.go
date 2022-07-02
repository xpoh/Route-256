package main

import (
	"bufio"
	. "fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	//var T int
	var n, q int
	var a []int

	//for Fscan(in, &T); T > 0; T-- {
	Fscan(in, &n, &q)
	a = make([]int, n)
	var t, id, msg, msg_all int

	for i := 0; i < q; i++ {
		Fscan(in, &t, &id)
		if t == 1 {
			msg++
			if id > 0 {
				a[id-1] = msg
			} else {
				msg_all = msg
			}
		} else {
			if msg_all > a[id-1] {
				Println(msg_all)
			} else {
				Println(a[id-1])
			}
		}
	}
	//}
}

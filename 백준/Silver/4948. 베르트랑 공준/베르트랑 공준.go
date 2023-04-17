package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func Prime(num int) int {
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return 0
		}
	}
	return 1
}

func main() {
	var n, cnt int

	s := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	for {
		s.Scan()
		n, _ = strconv.Atoi(s.Text())

		if n == 0 {
			break
		} else {
			cnt = 0
			for i := n + 1; i <= n*2; i++ {
				cnt += Prime(i)
			}
			fmt.Fprintln(w, cnt)
		}
	}
	w.Flush()
}
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n, fac int

	s := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	s.Scan()
	n, _ = strconv.Atoi(s.Text())

	if n == 0 {
		fmt.Fprintln(w, 1)
	} else {
		fac = 1
		for i := 1; i <= n; i++ {
			fac *= i
		}
		fmt.Fprintln(w, fac)
	}
	w.Flush()
}
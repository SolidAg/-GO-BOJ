package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n, money, qu, di, ni, pe int

	s := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	s.Scan()
	n, _ = strconv.Atoi(s.Text())

	for i := 0; i < n; i++ {
		s.Scan()
		money, _ = strconv.Atoi(s.Text())

		qu = money / 25
		money = money - (25 * qu)
		di = money / 10
		money = money - (10 * di)
		ni = money / 5
		pe = money - (5 * ni)

		fmt.Fprintf(w, "%d %d %d %d\n", qu, di, ni, pe)
	}

	w.Flush()
}
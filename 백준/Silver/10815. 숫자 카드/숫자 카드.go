package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int
	fmt.Fscanln(r, &n)

	var cards = map[int]int{}
	for i := 0; i < n; i++ {
		var input int
		if i == n-1 {
			fmt.Fscanln(r, &input)
		} else {
			fmt.Fscan(r, &input)
		}
		cards[input]++
	}

	var m int
	fmt.Fscanln(r, &m)

	for i := 0; i < m; i++ {
		var num int
		fmt.Fscan(r, &num)
		fmt.Fprintf(w, "%d ", hasCard(cards, num))
	}
}

func hasCard(cards map[int]int, num int) int {
	if cards[num] != 0 {
		return 1
	}
	return 0
}
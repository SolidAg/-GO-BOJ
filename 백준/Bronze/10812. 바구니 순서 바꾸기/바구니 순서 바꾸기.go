package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, m, i, j, k int

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	defer w.Flush()

	fmt.Fscanln(r, &n, &m)

	var basket []int = make([]int, n)
	var basket1 []int = make([]int, n)

	for p := 0; p < n; p++ {
		basket[p] = p + 1
		basket1[p] = p + 1
	}

	for p := 0; p < m; p++ {
		fmt.Fscanln(r, &i, &j, &k)

		c := i
		e := k

		for q := i; q <= i+j-k; q++ {
			basket[q-1] = basket1[e-1]
			e++
		}

		for q := i + j - k + 1; q <= j; q++ {
			basket[q-1] = basket1[c-1]
			c++
		}

		for p := 0; p < n; p++ {
			basket1[p] = basket[p]
		}
	}

	for p := 0; p < n; p++ {
		fmt.Fprintf(w, "%d ", basket[p])
	}
}
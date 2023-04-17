package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var m, n int

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	fmt.Fscanln(r, &m, &n)

	num := make([]bool, n+1)

	for i := 2; i < n+1; i++ {
		if num[i] {
			continue
		}
		if i >= m {
			fmt.Fprintln(w, i)
		}
		for j := i * 2; j <= n; j += i {
			num[j] = true
		}
	}

	w.Flush()
}
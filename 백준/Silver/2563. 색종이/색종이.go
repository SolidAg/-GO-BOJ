package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, x, y, count int

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	defer w.Flush()

	fmt.Fscanln(r, &n)
	var array [100][100]int

	for i := 0; i < n; i++ {
		fmt.Fscanf(r, "%d %d\n", &x, &y)
		for j := x; j < x+10; j++ {
			for k := y; k < y+10; k++ {
				array[j][k] += 1
			}
		}
	}

	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			if array[i][j] > 0 {
				count++
			}
		}
	}

	fmt.Fprintln(w, count)
}
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	defer w.Flush()

	fmt.Fscanln(r, &n)
	var a []int = make([]int, n)
	var b []int = make([]int, n)

	for i := range a {
		fmt.Fscanf(r, "%d %d\n", &a[i], &b[i])
		x := 1

		if a[i] > b[i] {
			for j := 1; j <= b[i]; j++ {
				if a[i]%j == 0 && b[i]%j == 0 {
					x = x * j
					a[i] = a[i] / j
					b[i] = b[i] / j
					j = 1
				}
			}
		} else if a[i] <= b[i] {
			for j := 1; j <= a[i]; j++ {
				if a[i]%j == 0 && b[i]%j == 0 {
					x = x * j
					a[i] = a[i] / j
					b[i] = b[i] / j
					j = 1
				}
			}
		}
		fmt.Fprintln(w, x*a[i]*b[i])
	}
}
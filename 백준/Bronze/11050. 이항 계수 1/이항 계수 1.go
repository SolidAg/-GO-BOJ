package main

import (
	"bufio"
	"fmt"
	"os"
)

func factorial(num int) int {
	var fac int
	fac = 1
	for i := 1; i <= num; i++ {
		fac *= i
	}
	return fac
}

func main() {
	var n, k int

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	fmt.Fscanln(r, &n, &k)

	fmt.Fprintln(w, factorial(n)/(factorial(n-k)*factorial(k)))

	w.Flush()
}
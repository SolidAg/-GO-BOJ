package main

import (
	"bufio"
	"fmt"
	"os"
)

func gcd(x, y int) int {
	if x < y {
		y, x = x, y
	}

	for y != 0 {
		x, y = y, x%y
	}

	return x
}

func main() {
	var a, b, c, d, m, g int

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	defer w.Flush()

	fmt.Fscanln(r, &a, &b)
	fmt.Fscanln(r, &c, &d)

	g = gcd(b, d)

	m = (b * d) / g
	a = (m / b) * a
	c = (m / d) * c
	a = a + c

	g = gcd(a, m)

	fmt.Fprintln(w, a/g, m/g)
}
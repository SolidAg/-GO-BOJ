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
	var a, b, g int

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	defer w.Flush()

	fmt.Fscanln(r, &a, &b)

	g = gcd(a, b)

	fmt.Fprintln(w, (a*b)/g)
}
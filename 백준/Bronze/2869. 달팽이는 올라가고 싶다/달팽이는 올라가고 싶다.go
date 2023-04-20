package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a, b, v int

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	fmt.Fscanln(r, &a, &b, &v)

	var day = 1

	if (v-a)%(a-b) == 0 {
		day += (v - a) / (a - b)
	} else {
		day += (v-a)/(a-b) + 1
	}

	fmt.Fprintln(w, day)
	w.Flush()
}
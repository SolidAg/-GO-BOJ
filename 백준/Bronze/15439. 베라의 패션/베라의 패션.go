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

	fmt.Fscanln(r, &n)
	fmt.Fprintln(w, n*(n-1))

	defer w.Flush()
}
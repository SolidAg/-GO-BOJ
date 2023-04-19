package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, bee int

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	fmt.Fscanln(r, &n)

	bee = 1
	for i := 0; ; i++ {

		bee = bee + 6*i
		if n > bee {
			continue
		} else {
			fmt.Fprintln(w, i+1)
			break
		}
	}
	w.Flush()
}
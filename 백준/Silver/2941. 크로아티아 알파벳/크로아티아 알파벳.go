package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var s string

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	defer w.Flush()

	fmt.Fscanln(r, &s)

	s = strings.Replace(s, "c=", "!", -1)
	s = strings.Replace(s, "c-", "@", -1)
	s = strings.Replace(s, "dz=", "#", -1)
	s = strings.Replace(s, "d-", "$", -1)
	s = strings.Replace(s, "lj", "%", -1)
	s = strings.Replace(s, "nj", "^", -1)
	s = strings.Replace(s, "s=", "&", -1)
	s = strings.Replace(s, "z=", "*", -1)

	fmt.Fprintln(w, len(s))
}
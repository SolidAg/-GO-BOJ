package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter((os.Stdout))

	s, _ := r.ReadString('\n')

	defer w.Flush()

	st := strings.Fields(s)

	fmt.Fprintln(w, len(st))
}
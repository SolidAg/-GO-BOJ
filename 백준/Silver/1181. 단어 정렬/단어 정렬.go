package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var n int

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	defer w.Flush()

	fmt.Fscanln(r, &n)
	var s []string = make([]string, n)

	for i := 0; i < n; i++ {
		fmt.Fscanln(r, &s[i])
	}

	sort.Slice(s, func(i, j int) bool {
		if len(s[i]) < len(s[j]) {
			return true
		} else if len(s[i]) == len(s[j]) {
			return s[i] < s[j]
		} else {
			return false
		}
	})

	for i := range s {
		if i > 0 && s[i-1] == s[i] {
			continue
		}
		fmt.Fprintln(w, s[i])
	}
}
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
	var xy = make([][]int, n)

	for i := 0; i < n; i++ {
		xy[i] = make([]int, 2)
		fmt.Fscanln(r, &xy[i][0], &xy[i][1])
	}

	sort.Slice(xy, func(i, j int) bool {
		for idx := range xy[i] {
			if xy[i][idx] == xy[j][idx] {
				continue
			}
			return xy[i][idx] < xy[j][idx]
		}
		return false
	})

	for i := 0; i < n; i++ {
		fmt.Fprintf(w, "%d %d\n", xy[i][0], xy[i][1])
	}
}
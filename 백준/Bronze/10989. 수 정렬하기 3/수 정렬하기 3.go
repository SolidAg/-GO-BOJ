package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var n int

	s := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	defer w.Flush()

	s.Scan()
	n, _ = strconv.Atoi(s.Text())

	var number []int = make([]int, n)

	for i := 0; i < n; i++ {
		s.Scan()
		number[i], _ = strconv.Atoi(s.Text())
	}

	sort.Slice(number, func(i, j int) bool {
		return number[i] < number[j]
	})

	for i := range number {
		fmt.Fprintln(w, number[i])
	}
}
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var n int

	s := bufio.NewScanner(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	s.Scan()
	n, _ = strconv.Atoi(s.Text())

	fmt.Fprintln(w, int(math.Sqrt(float64(n))))
	w.Flush()
}
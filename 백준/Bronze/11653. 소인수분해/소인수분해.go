package main

import (
	"bufio"
	_ "fmt"
	"os"
	"strconv"
	_ "strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	var N int
	if sc.Scan() {
		N, _ = strconv.Atoi(sc.Text())
	}
	if N == 1 {
		return
	}

	var r []int

	i := 2
	for i < N {
		if N%i == 0 {
			r = append(r, i)
			N /= i
			i = 2
		} else {
			i++
			if i*i > N {
				break
			}
		}
	}
	r = append(r, N)
	for _, v := range r {
		wr.WriteString(strconv.Itoa(v))
		wr.WriteString("\n")
	}
}

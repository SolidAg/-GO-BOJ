package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var t, g, sum, avg, count float32
	var s, grade string

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	defer w.Flush()

	for i := 0; i < 20; i++ {
		fmt.Fscanln(r, &s, &t, &grade)
		switch grade {
		case "A+":
			g = 4.5
			count += t
		case "A0":
			g = 4.0
			count += t
		case "B+":
			g = 3.5
			count += t
		case "B0":
			g = 3.0
			count += t
		case "C+":
			g = 2.5
			count += t
		case "C0":
			g = 2.0
			count += t
		case "D+":
			g = 1.5
			count += t
		case "D0":
			g = 1.0
			count += t
		case "F":
			g = 0.0
			count += t
		case "P":
			g = 0.0
		}
		sum += g * t
	}

	avg = sum / count
	fmt.Fprintln(w, avg)
}
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	var avg, max float32
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter((os.Stdout))

	defer w.Flush()

	fmt.Fscanln(r, &n)

	var score []float32 = make([]float32, n)

	for i := range score {
		fmt.Fscan(r, &score[i])

		if i == 0 {
			max = score[0]
		}

		if score[i] > max {
			max = score[i]
		}
	}

	for i := range score {
		score[i] = score[i] / max * 100
		avg += score[i]
	}

	fmt.Fprintln(w, avg/float32(n))
}
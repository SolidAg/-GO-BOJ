package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func change(array []string) {
	var temp string

	for i := range array {
		for j := i; j < len(array); j++ {
			if array[i] < array[j] {
				temp = array[i]
				array[i] = array[j]
				array[j] = temp
			}
		}
	}
}

func main() {
	var s string

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	defer w.Flush()

	fmt.Fscanln(r, &s)

	stirng := strings.Split(s, "")

	change(stirng)

	for i := range stirng {
		fmt.Fprintf(w, "%s", stirng[i])
	}
}
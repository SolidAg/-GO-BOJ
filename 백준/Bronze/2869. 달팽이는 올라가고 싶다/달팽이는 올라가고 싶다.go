package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	var a,b,v int
	r := bufio.NewReader(os.Stdin)
	fmt.Fscanln(r, &a,&b,&v)

	var day = 1

	if (v-a)%(a-b) == 0 {
		day += (v-a)/(a-b)
	} else {
		day += (v-a)/(a-b) + 1
	}

	fmt.Println(day)
}
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type userInfo struct {
	age  int
	name string
}

func main() {
	var n int
	r := bufio.NewReader(os.Stdin)
	fmt.Fscanln(r, &n)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var users = make([]userInfo, n)

	for i := 0; i < n; i++ {
		fmt.Fscanln(r, &users[i].age, &users[i].name)
	}

	sort.SliceStable(users, func(i, j int) bool {
		return users[i].age < users[j].age
	})

	for i := 0; i < n; i++ {
		fmt.Fprintln(w, users[i].age, users[i].name)
	}
}
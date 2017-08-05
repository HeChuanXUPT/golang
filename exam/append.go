package main

import (
	"fmt"
)

func Test_append() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)
}

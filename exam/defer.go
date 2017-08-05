package main

import (
	"fmt"
)

func Test_defer() {
	defer func() {
		fmt.Println("NO.1")
	}()
	defer func() {
		fmt.Println("NO.2")
	}()
	panic("die")
}

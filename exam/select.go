package main

import (
	"fmt"
	"runtime"
)

// tips:结果随机
/*
https://golang.org/ref/spec#Select_statements
If one or more of the communications can proceed,
a single one that can proceed is chosen via a uniform pseudo-random selection
*/

func Test_select() {
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	int_chan <- 3
	string_chan <- "hello"
	select {
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		fmt.Println(value)
	}
}

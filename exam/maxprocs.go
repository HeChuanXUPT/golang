package main

import (
	"fmt"
	"runtime"
	"sync"
)

/*
The GOMAXPROCS variable limits the number of operating system threads that
can execute user-level Go code simultaneously. There is no limit to the number
of threads that can be blocked in system calls on behalf of Go code; those do not
count against the GOMAXPROCS limit. This package's GOMAXPROCS function queries and changes the limit
*/

func Test_core() {
	runtime.GOMAXPROCS(1)
	sw := sync.WaitGroup{}
	max := 10
	sw.Add(max)

	for i := 0; i < max; i++ {
		go func() {
			sw.Done()
			fmt.Println(i)

		}()
	}

	sw.Wait()
}

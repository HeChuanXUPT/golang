package main

import (
	"fmt"
	"os"
)

func main() {

	Test_interface()
	os.Exit(0)
	fmt.Println("Test_inherit")
	Test_inherit()

	fmt.Println("Test_append")
	Test_append()
	fmt.Println("Test_select")
	Test_select()
	fmt.Println("Test_init")
	Test_init()
	fmt.Println("Test_core")
	Test_core()
	fmt.Println("Test_range")
	Test_range()
	fmt.Println("Test_defer")
	Test_defer()
}

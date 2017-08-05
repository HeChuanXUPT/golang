package main

import "fmt"

func Test_interface() {
	none := func(arg int) interface{} {
		var result *struct{} = nil
		if arg > 0 {
			result = &struct{}{}
		}
		return result
	}

	res := none(0)
	fmt.Println(res)
	res = none(1)
	fmt.Println(res)
}

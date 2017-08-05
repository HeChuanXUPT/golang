package main

import (
	"fmt"
)

// tips: range returns the key and the "copy" of the value

func Test_range() {
	m := make(map[string]*Player)
	ps := []Player{
		{Name: "hello", Money: 1234},
		{Name: "world", Money: 4321},
	}
	for _, e := range ps {
		m[e.Name] = &e
	}
	fmt.Println(m["hello"].Money)
	fmt.Println(m["world"].Money)
}

func Test_range2() {
	var ps [10]Player

	for i, e := range ps {
		e.Money = i
	}

	for _, e := range ps {
		fmt.Println(e.Money)
	}
}

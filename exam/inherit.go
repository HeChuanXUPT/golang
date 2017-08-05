package main

import (
	"fmt"
)

func init() {
	DefaultMoney = 0
}

var DefaultMoney = 10

type player struct {
	Money int
}

func (p *player) Update() {
	p.Money = DefaultMoney
}

func (p *player) ShowMoney() {
	p.Update()
	fmt.Println(p.Money)
}

type richPlayer struct {
	player
}

func (p *richPlayer) Update() {
	p.Money = 999
}

func Test_inherit() {
	p := richPlayer{}
	p.ShowMoney()
}

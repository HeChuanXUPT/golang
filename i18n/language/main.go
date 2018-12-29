package main

//go:generate gotext -srclang=en update -out=catalog/catalog.go -lang=en,zh

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
        _ "github.com/HeChuanXUP/golang/language/catalog"
)

func main() {
	p := message.NewPrinter(language.Chinese)
	p.Printf("Hello world!")
	p.Printf("\n")

}

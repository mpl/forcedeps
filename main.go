package main

import (
	"github.com/gopherjs/gopherjs/compiler"
	"myitcv.io/react"
)

func main() {
	foo := compiler.Version
	println(foo)
	var bar react.Element
	_ = bar
}

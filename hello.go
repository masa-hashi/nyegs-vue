package main

import (
	"fmt"

	"./api/hello"
)

func main() {
	s := hello.GetHello("HI! Guest! ")
	fmt.Println(s)
}

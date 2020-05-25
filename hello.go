package main

import (
    "fmt"
    "./hello"
)
  
func main() {
  s := hello.GetHello("HI! Guest! ")
  fmt.Println(s)
}
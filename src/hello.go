package main

import "fmt"

func main() {
	var (
		hello, name string // объявление нескольких переменных
		punct       = "!"  // не явная типизация
	)
	name = "Go"
	hello = "Hello " + name + punct
	fmt.Println(hello)
}

package main

import (
	"fmt"
	"reflect"
)

func show_var(variable string, value reflect.Value) {
	fmt.Printf("%s is %v type %s\n", variable, value, value.Kind())
}

func main() {
	var a int = 32
	show_var("a", reflect.ValueOf(a))

	b := 16.8
	show_var("b", reflect.ValueOf(b))

	var a_plus_b = a + int(b)
	show_var("a_plus_b", reflect.ValueOf(a_plus_b))

	var b_minus_a = int(b) - a
	show_var("a_minus_b", reflect.ValueOf(b_minus_a))

	var a_multiple_b = float64(a) * b
	show_var("a_multiple_b", reflect.ValueOf(a_multiple_b))

	var b_div_a = b / float64(a)
	show_var("b_div_a", reflect.ValueOf(b_div_a))

	var b_rem_div_a = int(b) % a
	show_var("b_rem_div_a", reflect.ValueOf(b_rem_div_a))

	// Постфиксный инкремент a++
	// Постфиксный декремент a--
}

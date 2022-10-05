package main

import "fmt"

func main() {
	var result = 20 + 20
	fmt.Println(result)

	var a = 10
	var b = 20
	var c = 30
	var d = a + b*c
	fmt.Println(d)

	// Augmented Assignments

	var i = 10
	i += 10 // i = i + 10
	fmt.Println(i)

	// Unary Operator

	i++
	fmt.Println(i)

	var negative = -100
	var positive = +100
	fmt.Println(negative)
	fmt.Println(positive)
}

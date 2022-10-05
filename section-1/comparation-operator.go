package main

import "fmt"

func main() {
	var name1 = "Fajar"
	var name2 = "Tiara"

	var result bool = name1 == name2
	fmt.Println(result)

	var value1 = 200
	var value2 = 300
	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}

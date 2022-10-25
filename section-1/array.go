package main

import "fmt"

func main() {

	// membuat array manual
	var names [3]string

	names[0] = "Fajar"
	names[1] = "Dwi"
	names[2] = "Rianto"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	//	membuat array langsung
	var values = [3]int{
		90,
		91,
		92,
	}

	fmt.Println(values)
	fmt.Println(values[0])

	fmt.Println(len(names))
	fmt.Println(len(values))

	var lagi [10]string
	fmt.Println(len(lagi))
}

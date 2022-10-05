package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var ktpFajar NoKTP = "111111111"
	var marriedStatus Married = true
	fmt.Println(ktpFajar)
	fmt.Println(marriedStatus)
}

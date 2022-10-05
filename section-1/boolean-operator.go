package main

import "fmt"

func main() {
	var nilaiAkhir = 90
	var aBsensi = 60

	var lulusUjian = nilaiAkhir > 80
	var lulusAbsensi = aBsensi >= 80
	fmt.Println(lulusAbsensi)
	fmt.Println(lulusUjian)

	var lulus = lulusAbsensi && lulusUjian
	fmt.Println(lulus)

	fmt.Println(nilaiAkhir >= 80 && aBsensi >= 80)
}

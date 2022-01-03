package main

import "fmt"

func main() {
	var nilaiakhir = 90
	var absensi = 80

	var lulusnilaiakhir bool = nilaiakhir > 80
	var lulusabsensi bool = absensi > 80

	var lulus bool = lulusnilaiakhir && lulusabsensi

	fmt.Println(lulus)

	fmt.Println(nilaiakhir >= 80 && absensi >= 80)

}

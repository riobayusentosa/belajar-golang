package main

import "fmt"

func main() {
	type nik string
	var ktpRio nik = "111111111111"

	fmt.Println(ktpRio)
	fmt.Println(nik("2222222222"))
}

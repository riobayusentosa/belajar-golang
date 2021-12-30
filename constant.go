package main

import "fmt"

func main() {
	const firstname string = "Rio"
	const lastname = "Bayu"

	const value = 1000
	fmt.Println(value)

	const (
		vfirstname = "Rio"
		vlastname  = "Bayu"
		vvalue     = 100
	)

	fmt.Println(vfirstname)

}

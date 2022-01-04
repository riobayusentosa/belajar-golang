package main

import "fmt"

func main() {
	name := "Rio"

	if name == "Rio " {
		fmt.Println("Hello", name)
	} else if name == "Rio" {
		fmt.Println("Ini dia")
	} else {
		fmt.Println("Hi, Admin")
	}

	if length := len(name); length > 5 {
		fmt.Println("Nama terlalu panjang")
	} else {
		fmt.Println("Cocok")
	}
}

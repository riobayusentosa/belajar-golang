package main

import "fmt"

func main() {
	var name [3]string
	name[0] = "Rio"
	name[1] = "Bayu"
	name[2] = "Sentosa"

	fmt.Println(name[0])
	fmt.Println(name[1])
	fmt.Println(name[2])

	var values = [3]int{
		90, 80, 75,
	}

	fmt.Println(values)
	fmt.Println(values[0])

	fmt.Println(len(name))
	fmt.Println(len(values))
}

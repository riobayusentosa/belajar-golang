package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Rio",
		"address": "padang",
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	book := make(map[string]string)
	book["title"] = "Buku Go-Lang"
	book["author"] = "Penulis"
	book["wrong"] = "Ups"
	fmt.Println(book)
	fmt.Println(len(book))
	delete(book, "wrong")

	fmt.Println(book)
	fmt.Println(len(book))
}

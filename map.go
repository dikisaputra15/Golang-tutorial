package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "diki",
		"address": "kembang",
	}

	person["title"] = "programmer"
	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])

	var book map[string]string = make(map[string]string)
	book["title"] = "belajar golang"
	book["author"] = "diki"
	book["ups"] = "salah"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")
	fmt.Println(book)
	fmt.Println(len(book))
}

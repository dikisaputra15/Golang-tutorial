package main

import "fmt"

func main() {
	name := "shintaehyoungkimpangbon"

	switch name {
	case "eko":
		fmt.Println("hai eko")
		fmt.Println("hai eko")
	case "joko":
		fmt.Println("hai joko")
		fmt.Println("hai joko")
	default:
		fmt.Println("kenalan dong")
		fmt.Println("kenalan dong")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("nama terlalu panjang")
	case false:
		fmt.Println("nama sudah benar")
	}

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("nama sangat panjang")
	case length > 5:
		fmt.Print("nama lumayan panjang")
	default:
		fmt.Println("nama sudah benar")
	}
}

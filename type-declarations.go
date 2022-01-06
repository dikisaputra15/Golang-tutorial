package main

import "fmt"

func main() {
	type noktp string
	type married bool

	var noktpdiki noktp = "16009305930439"
	var marriedstatus married = true

	fmt.Println(noktpdiki)
	fmt.Println(marriedstatus)
}

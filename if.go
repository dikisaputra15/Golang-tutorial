package main

import "fmt"

func main() {
	var nama = "kurniawan"

	if nama == "eko" {
		fmt.Println("hai nama saya eko")
	} else if nama == "joko" {
		fmt.Println("saya joko")
	} else if nama == "budi" {
		fmt.Println("saya budi")
	} else {
		fmt.Println("hai kenalan dong")
	}

	if length := len(nama); length > 5 {
		fmt.Println("nama terlalu panjang")
	} else {
		fmt.Println("nama sudah benar")
	}
}

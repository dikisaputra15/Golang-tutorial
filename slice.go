package main

import "fmt"

func main() {
	var month = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = month[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	//month[5] = "Diubah"
	//fmt.Println(slice1)

	//slice1[0] = "mei lagi"
	//fmt.Println(month)

	var slice2 = month[10:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "eko")
	fmt.Println(slice3)
	slice3[1] = "bukan desember"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(month)

	newslice := make([]string, 2, 5)
	newslice[0] = "eko"
	newslice[1] = "kurniawan"
	fmt.Println(newslice)
	fmt.Println(len(newslice))
	fmt.Println(cap(newslice))

	copyslice := make([]string, len(newslice), cap(newslice))
	copy(copyslice, newslice)
	fmt.Println(copyslice)

	iniaray := [5]int{1, 2, 3, 4, 5}
	inislice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniaray)
	fmt.Println(inislice)
}

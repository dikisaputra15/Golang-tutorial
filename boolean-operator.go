package main

import "fmt"

func main() {
	var ujian = 80
	var absen = 80

	var lulusujian = ujian > 80
	var lulusabsen = absen > 80
	fmt.Println(lulusujian)
	fmt.Println(lulusabsen)

	var lulus = lulusujian && lulusabsen
	fmt.Println(lulus)
}

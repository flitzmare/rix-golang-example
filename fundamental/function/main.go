package main

import "fmt"

func main() {
	hasil := tambah(1, 2)
	fmt.Println("Hasilnya:", hasil)

	hasilPembagian, sisaPembagian := bagi(9, 3)
	fmt.Println("Hasilnya:", hasilPembagian)
	fmt.Println("Sisanya:", sisaPembagian)

	hasilPembagian, sisaPembagian = bagi(9, 3)
	fmt.Println("Hasilnya:", hasilPembagian)
	fmt.Println("Sisanya:", sisaPembagian)
}

func hello() {
	fmt.Println("Hello")
}

func tambah(a, b int) int {
	hasil := a + b
	return hasil
}

func bagi(a, b int) (int, int) {
	hasil := a / b
	sisa := a % b
	return hasil, sisa
}

package main

import "fmt"

func main() {
	var score = 60
	if score >= 90 {
		fmt.Println("Nilai anda A")
	} else if score >= 80 {
		fmt.Println("Nilai anda B")
	} else if score >= 70 {
		fmt.Println("Nilai anda C")
	} else {
		fmt.Println("Kamu tidak lulus!")
	}
}

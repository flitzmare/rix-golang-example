package main

import "fmt"

func main() {
	var point = 7
	switch point {
	case 10:
		fmt.Println("Sempurna")
	case 7,8,9:
		fmt.Println("Luar Biasa!")
	default:
		fmt.Println("Okay!")
	}
}

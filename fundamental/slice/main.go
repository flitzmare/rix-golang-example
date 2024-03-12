package main

import "fmt"

func main() {
	var buah = []string{"pisang", "apel", "melon", "mangga"}

	var subSlice = buah[1:2]

	fmt.Println(subSlice)
}

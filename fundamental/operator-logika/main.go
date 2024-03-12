package main

import "fmt"

func main() {
	var a = false
	var b = true

	var aANDb = a && b
	fmt.Println(aANDb)

	var aORb = a || b
	fmt.Println(aORb)

	var notA = !a
	fmt.Println(notA)
}

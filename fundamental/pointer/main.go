package main

import "fmt"

func main() {
	var name = "Riksa"
	var secondName *string = &name

	name = "Suviana"

	fmt.Println(name)
	fmt.Println(*secondName)
}

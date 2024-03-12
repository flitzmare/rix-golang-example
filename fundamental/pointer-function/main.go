package main

import "fmt"

func main() {
	var name string = "Riksa"
	SetName(&name)
	fmt.Println(name)

	name = SetNameWithoutPointer()
	fmt.Println(name)
}

func SetName(name *string) {
	*name = "Suviana"
}

func SetNameWithoutPointer() string {
	return "Suviana"
}

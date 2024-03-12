package main

import (
	"fmt"
	"strconv"
)

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Ada panic: ", r)
		} else {
			fmt.Println("Tidak ada panic error!")
		}
	}()

	_, err := strconv.Atoi("asd")
	if err != nil {
		panic(err)
	}

	fmt.Println("Kode ini masih berlanjut")
}

package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i == 3 {
			continue
		}
		fmt.Println("Pengulangan ke-", i)
	}

	var i = 1
	for {
		fmt.Println("Angka", i)
		i++

		if i == 6 {
			break
		}
	}

	var buah = []string{"pisang", "apel", "mangga"}
	for index, item := range buah {
		fmt.Println("Index", index, item)
	}
}

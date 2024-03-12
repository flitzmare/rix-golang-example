package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World!")
	})
	http.HandleFunc("/template", func(w http.ResponseWriter, r *http.Request) {
		var data = map[string]string{
			"Judul": "Belajar Golang",
			"Nama":  "Riksa",
			"Pesan": "Halo, selamat datang!",
		}

		t, err := template.ParseFiles("hello.html")
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		t.Execute(w, data)
	})
	http.ListenAndServe(":8080", nil)
}

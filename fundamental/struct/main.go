package main

import "fmt"

func main() {
	dompetSaya := Dompet{
		Uang:  100000,
		Kartu: []string{"Bank A", "KTP"},
	}

	dompetSaya.Uang = 50000

	dompetSaya.Kartu = append(dompetSaya.Kartu, "SIM C")

	tasSaya := Tas{
		Pemilik:       "Riksa",
		DompetPemilik: dompetSaya,
	}

	fmt.Println(tasSaya.DompetPemilik.Uang)
}

type Dompet struct {
	Uang  float64
	Kartu []string
}

type Tas struct {
	Pemilik       string
	DompetPemilik Dompet
}

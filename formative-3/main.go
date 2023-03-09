package main

import "fmt"

func main() {
	// soal 1
	var panjangPersegiPanjang int = 8
	var lebarPersegiPanjang int = 5

	var alasSegitiga float32 = 6.0
	var tinggiSegitiga float32 = 7.0

	var luasPersegiPanjang int
	var kelilingPersegiPanjang int
	var luasSegitiga float32

	luasPersegiPanjang = panjangPersegiPanjang * lebarPersegiPanjang
	luasSegitiga = 0.5 * alasSegitiga * tinggiSegitiga
	kelilingPersegiPanjang = 2 * (panjangPersegiPanjang + lebarPersegiPanjang)

	fmt.Println(luasPersegiPanjang)
	fmt.Println(luasSegitiga)
	fmt.Println(kelilingPersegiPanjang)

	// soal 2
	var nilaiJohn = 80
	var nilaiDoe = 50

	var nilai int

	nilai = nilaiJohn | nilaiDoe
	var indeks string

	if nilai >= 80 {
		indeks = "A"
	} else if nilai >= 70 && nilai < 80 {
		indeks = "B"
	} else if nilai >= 60 && nilai < 70 {
		indeks = "C"
	} else if nilai >= 50 && nilai < 60 {
		indeks = "D"
	} else {
		indeks = "E"
	}

	fmt.Println(indeks)

	//soal 3
	var tanggal = 17
	var bulan = 8
	var tahun = 1945

	switch {
	case tanggal == 17:
		tanggal = 5
	case bulan == 8:
		bulan = 7
	case tahun == 1945:
		tahun = 2000
	}

	fmt.Println(tanggal, bulan, tahun)

	// soal 4
	if tahun == 1964 && tahun < 1965 {
		fmt.Println("Baby boomer")
	} else if tahun == 1965 && tahun < 1980 {
		fmt.Println("Generasi X")
	} else if tahun == 1980 && tahun < 1995 {
		fmt.Println("Generasi Y (Millenials)")
	} else if tahun == 1995 && tahun < 2016 {
		fmt.Println("Generasi Z")
	}
}

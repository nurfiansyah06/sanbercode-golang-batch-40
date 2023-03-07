package main

import (
	"fmt"
	"strings"
)

func main() {
	// soal 1
	// jawaban 1
	fmt.Println("Bootcamp Digital Skill Sanbercode Golang")

	// soal 2
	// jawaban 2
	var str = "Halo Dunia"
	fmt.Println(strings.Replace(str, "Halo Dunia", "Halo Golang", -1))

	// soal 3
	// jawaban 3
	var kataPertama = "saya"
	var kataKedua = "senang"
	var kataKetiga = "belajar"
	var kataKeempat = "golang"

	fmt.Println(kataPertama + " " + kataKedua + " " + kataKetiga + " " + kataKeempat)

	// soal 4
	// jawaban 4
	var angkaPertama = 8
	var angkaKedua = 5
	var angkaKetiga = 6
	var angkaKeempat = 7
	result := angkaPertama + angkaKedua + angkaKetiga + angkaKeempat

	fmt.Println(result)

	// soal 5
	// jawaban 5
	kalimat := "halo halo bandung"
	angka := 2021

	fmt.Println(kalimat, "-", angka)
}

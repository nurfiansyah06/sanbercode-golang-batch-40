package main

import "fmt"

func main() {
	// soal 1
	for i := 1; i < 21; i++ {
		if i%2 == 0 {
			fmt.Println(i, " Berkualitas")
		} else if i%2 != 0 {
			fmt.Println(i, " Santai")
		} else if i%3 == 0 {
			fmt.Println(i, " I Love Coding")
		}
	}

	//soal 2

	//soal 3
	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
	fmt.Println(kalimat)

	//soal 4
	var sayuran = [...]string{"Bayam", "Buncis", "Kangkung", "Kubis", "Seledri", "Tauge", "Timun"}
	for i := 1; i < len(sayuran); i++ {
		fmt.Println(i, sayuran[i])
	}

	//soal 5
	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}

	resultVolume := satuan["panjang"] * satuan["lebar"] * satuan["tinggi"]
	for i, v := range satuan {
		fmt.Println(i, v)
	}

	fmt.Println("Volume Balok:", resultVolume)
}

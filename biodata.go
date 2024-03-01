package main

import (
	"fmt"
	"os"
)

type Teman struct {
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

var dataTeman = map[int]Teman{
	1: {"Mamat", "Jakarta", "Mahasiswa", "Golang itu mudah"},
	2: {"Asep", "Palembang", "Freelance", "Golang itu gak ribet"},
	3: {"Roni", "Pekanbaru", "Fresh Graduate", "Golang itu enak banget"},
}

// menampilkan informasi teman sesuai nomor absen
func TampilkanDataTeman(absen int) {
	teman, ok := dataTeman[absen]
	if !ok {
		fmt.Println("gak ada datanya cuy.")
		return
	}

	fmt.Printf("Data absen %d:\n", absen)
	fmt.Printf("Nama: %s\n", teman.nama)
	fmt.Printf("Alamat: %s\n", teman.alamat)
	fmt.Printf("Pekerjaan: %s\n", teman.pekerjaan)
	fmt.Printf("Alasan memilih kelas Golang: %s\n", teman.alasan)
}

func main() {
	// ngechek argumen
	if len(os.Args) != 2 {
		fmt.Println("Hint: go run biodata.go *nomor")
		return
	}

	absen := os.Args[1]

	// Konversi no absen ke integer
	var absenInt int
	_, err := fmt.Sscanf(absen, "%d", &absenInt)
	if err != nil {
		fmt.Println("gak bisa huruf cuy, harus angka.")
		return
	}

	TampilkanDataTeman(absenInt)
}

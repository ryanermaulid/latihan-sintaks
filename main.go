package main

import "fmt"

func main() {

	// Variabel & Struktur Data

	fmt.Println("=== Variabel & Struktur Data ===")

	// 1. Deklarasi 5 variabel dengan tipe berbeda
	var nama string = "Ryan"
	var umur int = 21
	var ipk float64 = 3.85
	var isActive bool = true
	hobi := []string{"Nonton Film", "Gaming"}

	fmt.Printf("Nama: %s (Tipe: %T)\n", nama, nama)
	fmt.Printf("Umur: %d (Tipe: %T)\n", umur, umur)
	fmt.Printf("IPK: %.2f (Tipe: %T)\n", ipk, ipk)
	fmt.Printf("Status Aktif: %v (Tipe: %T)\n", isActive, isActive)
	fmt.Printf("Hobi: %v (Tipe: %T)\n\n", hobi, hobi)

	// 2. Map data mahasiswa (nama -> nilai)
	nilaiMahasiswa := make(map[string]int)

	// a. Menambah data
	nilaiMahasiswa["Ryan"] = 90
	nilaiMahasiswa["Vano"] = 85
	nilaiMahasiswa["Cipmank"] = 78

	// b. Membaca dengan pengecekan keberadaan (comma ok)
	if val, ada := nilaiMahasiswa["Cipmank"]; ada {
		fmt.Printf("Nilai Cipmank: %d\n", val)
	}

	// c. Menghapus data
	delete(nilaiMahasiswa, "Vano")

	// d. Menelusuri seluruh data map (range)
	fmt.Println("Daftar nilai mahasiswa setelah update:")
	for k, v := range nilaiMahasiswa {
		fmt.Printf("- %s: %d\n", k, v)
	}
}
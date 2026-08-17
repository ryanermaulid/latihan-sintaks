package main

import "fmt"

// Menukar nilai dua integer melalui pointer
func swap(a, b *int) {
	*a, *b = *b, *a
}

// Menambahkan item baru ke slice lewat pointer
func updateSlice(s *[]string, newItem string) {
	*s = append(*s, newItem)
}

// Pembanding pass by value vs pointer
func ubahValue(x int) {
	x = 999
}

func ubahPointer(x *int) {
	*x = 999
}

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

	// pointer
	
	fmt.Println("\n=== Pointer ===")

	// a. Demo swap
	x, y := 10, 20
	fmt.Printf("Sebelum swap: x = %d, y = %d\n", x, y)
	swap(&x, &y)
	fmt.Printf("Setelah swap : x = %d, y = %d\n", x, y)

	// b. Demo updateSlice
	listBuah := []string{"Apel", "Mangga"}
	fmt.Println("Sebelum updateSlice:", listBuah)
	updateSlice(&listBuah, "Jeruk")
	fmt.Println("Setelah updateSlice :", listBuah)

	// c. Perbandingan Pass by Value vs Pass by Pointer
	angka := 50
	ubahValue(angka)
	fmt.Println("Setelah pass by value  :", angka) // tetap 50
	ubahPointer(&angka)
	fmt.Println("Setelah pass by pointer:", angka) // berubah jadi 999
}
package main

import (
	"fmt"
)

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

func permutasi(n, r int) int {
	return factorial(n) / factorial(n-r)
}

func kombinasi(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func main() {
	var a, b, c, d int

	// Input dari pengguna
	fmt.Print("Masukkan empat bilangan (a b c d): ")
	fmt.Scan(&a, &b, &c, &d)

	// Hitung permutasi dan kombinasi
	hasilPertama := [2]int{permutasi(a, c), kombinasi(a, c)}
	hasilKedua := [2]int{permutasi(b, d), kombinasi(b, d)}

	// Tampilkan hasil
	fmt.Printf("%d %d\n", hasilPertama[0], hasilPertama[1])
	fmt.Printf("%d %d\n", hasilKedua[0], hasilKedua[1])
}

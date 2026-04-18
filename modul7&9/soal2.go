package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := range arr {
		fmt.Scan(&arr[i])
	}

	// a. semua
	fmt.Println(arr)

	// b. indeks ganjil
	for i := 1; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	// c. indeks genap
	for i := 0; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

	// d. kelipatan x
	var x int
	fmt.Scan(&x)
	for i := 0; i < len(arr); i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

	// e. hapus indeks
	var idx int
	fmt.Scan(&idx)
	arr = append(arr[:idx], arr[idx+1:]...)
	fmt.Println(arr)

	// f. rata-rata
	total := 0
	for _, v := range arr {
		total += v
	}
	rata := float64(total) / float64(len(arr))
	fmt.Printf("%.2f\n", rata)

	// g. standar deviasi
	variance := 0.0
	for _, v := range arr {
		d := float64(v) - rata
		variance += d * d
	}
	fmt.Printf("%.2f\n", math.Sqrt(variance/float64(len(arr))))

	// h. frekuensi
	var cari int
	fmt.Scan(&cari)
	count := 0
	for _, v := range arr {
		if v == cari {
			count++
		}
	}
	fmt.Println(count)
}
package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	// array kapasitas 1000
	var berat [1000]float64

	// input data
	for i := 0; i < N; i++ {
		fmt.Scan(&berat[i])
	}

	// asumsi awal
	min := berat[0]
	max := berat[0]

	// cari min & max
	for i := 1; i < N; i++ {
		if berat[i] < min {
			min = berat[i]
		}
		if berat[i] > max {
			max = berat[i]
		}
	}

	// output
	fmt.Println("Berat terkecil:", min)
	fmt.Println("Berat terbesar:", max)
}
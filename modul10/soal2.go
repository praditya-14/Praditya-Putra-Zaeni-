package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	var ikan [1000]float64

	// input berat ikan
	for i := 0; i < x; i++ {
		fmt.Scan(&ikan[i])
	}

	var hasil []float64

	// proses per wadah
	for i := 0; i < x; i += y {
		sum := 0.0

		for j := i; j < i+y && j < x; j++ {
			sum += ikan[j]
		}

		hasil = append(hasil, sum)
	}

	// output total tiap wadah
	for i := 0; i < len(hasil); i++ {
		fmt.Print(hasil[i], " ")
	}
	fmt.Println()

	// hitung rata-rata
	total := 0.0
	for i := 0; i < len(hasil); i++ {
		total += hasil[i]
	}

	rata := total / float64(len(hasil))
	fmt.Println(rata)
}
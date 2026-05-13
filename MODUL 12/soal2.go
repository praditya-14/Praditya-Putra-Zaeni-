package main

import "fmt"

func main() {
	var x, total, sah int
	suara := [21]int{}

	for {
		fmt.Scan(&x)
		if x == 0 {
			break
		}

		total++

		if x >= 1 && x <= 20 {
			sah++
			suara[x]++
		}
	}

	ketua, wakil := 1, 1

	for i := 2; i <= 20; i++ {
		if suara[i] > suara[ketua] {
			wakil = ketua
			ketua = i
		} else if i != ketua && suara[i] > suara[wakil] {
			wakil = i
		}
	}

	fmt.Println("Suara masuk:", total)
	fmt.Println("Suara sah:", sah)
	fmt.Println("Ketua RT:", ketua)
	fmt.Println("Wakil ketua:", wakil)
}
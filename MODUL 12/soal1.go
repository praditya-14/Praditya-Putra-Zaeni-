package main

import "fmt"

func main() {
	var x int
	total := 0
	sah := 0

	suara := make([]int, 21)

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

	fmt.Printf("Suara masuk: %d\n", total)
	fmt.Printf("Suara sah: %d\n", sah)

	for i := 1; i <= 20; i++ {
		if suara[i] > 0 {
			fmt.Printf("%d: %d\n", i, suara[i])
		}
	}
}
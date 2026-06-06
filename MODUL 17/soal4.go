package main

import "fmt"

func main() {
	var n int
	var S float64

	fmt.Print("N suku pertama: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		if i%2 == 1 {
			S += 1.0 / float64(2*i-1)
		} else {
			S -= 1.0 / float64(2*i-1)
		}
	}

	pi := 4 * S

	fmt.Printf("Hasil PI: %.7f\n", pi)
}
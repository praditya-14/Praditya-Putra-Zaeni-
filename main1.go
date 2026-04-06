package main

import "fmt"

// procedure faktorial
func faktorial(n int) int {
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

// procedure permutasi P(n, r)
func permutasi(n, r int) int {
	return faktorial(n) / faktorial(n-r)
}

// procedure kombinasi C(n, r)
func kombinasi(n, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}

func main() {
	var a, b, c, d int

	// input
	fmt.Scan(&a, &b, &c, &d)

	// baris pertama (a terhadap c)
	fmt.Println(permutasi(a, c), kombinasi(a, c))

	// baris kedua (b terhadap d)
	fmt.Println(permutasi(b, d), kombinasi(b, d))
}
package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	
	fmt.Print("N suku pertama: ")
	fmt.Scanln(&n) 

	sPrev := 4.0
	var sCurr float64

	for i := 2; i <= n; i++ {
		term := math.Pow(-1, float64(i-1)) * (4.0 / (2.0*float64(i) - 1.0))
		sCurr = sPrev + term

		if math.Abs(sCurr-sPrev) <= 0.00001 {
			fmt.Printf("Hasil PI: %.10f\n", sPrev)
			fmt.Printf("Hasil PI: %.10f\n", sCurr)
			fmt.Printf("Pada i ke: %d\n", i+1)
			break
		}
		sPrev = sCurr
	}
}
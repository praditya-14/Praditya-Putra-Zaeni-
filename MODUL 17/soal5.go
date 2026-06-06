package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(42) 

	var banyakTopping int
	fmt.Print("Banyak Topping: ")
	fmt.Scan(&banyakTopping)

	xc, yc, r := 0.5, 0.5, 0.5
	toppingPadaPizza := 0

	for i := 0; i < banyakTopping; i++ {
		x := rand.Float64()
		y := rand.Float64()

		if (x-xc)*(x-xc)+(y-yc)*(y-yc) <= r*r {
			toppingPadaPizza++
		}
	}

	pi := 4.0 * float64(toppingPadaPizza) / float64(banyakTopping)

	fmt.Printf("Topping pada Pizza: %d\n", toppingPadaPizza)
	fmt.Printf("PI : %.10f\n", pi)
}
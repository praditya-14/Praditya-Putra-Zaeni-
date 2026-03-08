package main

import "fmt"

func main() {
	var N int
	var bunga string
	pita := ""
	jumlah := 0

	fmt.Print("N : ")
	fmt.Scanln(&N)

	if N > 0 {

		for i := 1; i <= N; i++ {
			fmt.Printf("Bunga %d: ", i)
			fmt.Scanln(&bunga)

			pita = pita + bunga + " - "
			jumlah++
		}

	} else {

		i := 1
		for {
			fmt.Printf("Bunga %d: ", i)
			fmt.Scanln(&bunga)

			if bunga == "SELESAI" {
				break
			}

			pita = pita + bunga + " - "
			jumlah++
			i++
		}

	}

	fmt.Println("Pita:", pita)
	fmt.Println("Bunga:", jumlah)
}
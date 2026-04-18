package main

import "fmt"

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var hasil []string

	// input nama klub
	fmt.Print("Klub A: ")
	fmt.Scan(&klubA)

	fmt.Print("Klub B: ")
	fmt.Scan(&klubB)

	no := 1

	for {
		fmt.Printf("Pertandingan %d: ", no)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			fmt.Println("Hasil", no, ":", klubA)
			hasil = append(hasil, klubA)
		} else if skorB > skorA {
			fmt.Println("Hasil", no, ":", klubB)
			hasil = append(hasil, klubB)
		} else {
			fmt.Println("Hasil", no, ": Draw")
		}

		no++
	}

	fmt.Println("Pertandingan selesai")
}
package main

import "fmt"

type Domino struct {
	kiri, kanan int
}

func tampilKartu(kartu []Domino) {
	for i, k := range kartu {
		fmt.Printf("%d. [%d|%d]\n", i+1, k.kiri, k.kanan)
	}
}

func main() {
	// Kartu pemain
	kartu := []Domino{
		{2, 5},
		{5, 3},
		{3, 6},
		{6, 1},
	}

	// Kartu awal di meja
	rantai := []Domino{{1, 2}}

	fmt.Println("=== PERMAINAN GAPLEH ===")

	for len(kartu) > 0 {

		fmt.Print("\nRantai: ")
		for _, k := range rantai {
			fmt.Printf("[%d|%d] ", k.kiri, k.kanan)
		}

		fmt.Println("\n\nKartu Anda:")
		tampilKartu(kartu)

		ujung := rantai[len(rantai)-1].kanan

		fmt.Printf("\nAngka yang harus disambung: %d\n", ujung)

		var pilih int
		fmt.Print("Pilih nomor kartu: ")
		fmt.Scan(&pilih)

		if pilih < 1 || pilih > len(kartu) {
			fmt.Println("Pilihan tidak valid!")
			continue
		}

		k := kartu[pilih-1]

		if k.kiri == ujung {

			rantai = append(rantai, k)

			kartu = append(kartu[:pilih-1], kartu[pilih:]...)

			fmt.Println("Kartu berhasil dimainkan!")

		} else {

			fmt.Println("Kartu tidak bisa dimainkan!")
		}
	}

	fmt.Println("\n====================")
	fmt.Println("SELAMAT ANDA MENANG!")
	fmt.Println("====================")
}
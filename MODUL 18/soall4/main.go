package main

import "fmt"

var data string
var pos int
var currentChar byte

// start()
func start() {
	pos = 0
	currentChar = data[pos]
}

// maju()
func maju() {
	pos++
	if pos < len(data) {
		currentChar = data[pos]
	}
}

// eop()
func eop() bool {
	return currentChar == '.'
}

// cc()
func cc() byte {
	return currentChar
}

func main() {
	fmt.Print("Masukkan kalimat (akhiri dengan titik): ")
	fmt.Scanln(&data)

	start()

	jumlahKarakter := 0
	jumlahA := 0
	jumlahLE := 0

	var prev byte

	fmt.Println("\nKarakter yang terbaca:")

	for !eop() {
		fmt.Printf("%c ", cc())

		jumlahKarakter++

		if cc() == 'A' {
			jumlahA++
		}

		if prev == 'L' && cc() == 'E' {
			jumlahLE++
		}

		prev = cc()
		maju()
	}

	fmt.Println("\n\nHasil:")
	fmt.Println("Jumlah karakter =", jumlahKarakter)
	fmt.Println("Jumlah huruf A =", jumlahA)

	if jumlahKarakter > 0 {
		frekuensi := float64(jumlahA) / float64(jumlahKarakter)
		fmt.Println("Frekuensi A =", frekuensi)
	}

	fmt.Println("Jumlah pasangan LE =", jumlahLE)
}
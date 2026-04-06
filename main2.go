package main

import (
	"bufio"
	"fmt"
	"os"
)

// procedure hitungSkor
func hitungSkor(soal *int, skor *int, times []int) {
	*soal = 0
	*skor = 0

	for _, t := range times {
		if t < 301 { // dianggap berhasil kalau < 301
			*soal++
			*skor += t
		}
	}
}

func main () {
	scanner := bufio.NewScanner(os.Stdin)

	var pemenang string
	maxSoal := -1
	minWaktu := 1<<31 - 1

	for {
		scanner.Scan()
		line := scanner.Text()

		if line == "Selesai" {
			break
		}

		var nama string
		var t [8]int

		fmt.Sscanf(line, "%s %d %d %d %d %d %d %d %d",
			&nama, &t[0], &t[1], &t[2], &t[3],
			&t[4], &t[5], &t[6], &t[7])

		var soal, skor int
		hitungSkor(&soal, &skor, t[:])

		// tentukan pemenang
		if soal > maxSoal || (soal == maxSoal && skor < minWaktu) {
			pemenang = nama
			maxSoal = soal
			minWaktu = skor
		}
	}

	fmt.Println(pemenang, maxSoal, minWaktu)
}
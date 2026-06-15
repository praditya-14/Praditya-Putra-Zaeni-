package main

import "fmt"

type Domino struct {
	kiri, kanan int
}


func galiKartu(kartu []Domino, acuan Domino) {
	for _, k := range kartu {
		if k.kiri == acuan.kiri || k.kiri == acuan.kanan ||
			k.kanan == acuan.kiri || k.kanan == acuan.kanan {
			fmt.Println("Kartu cocok:", k)
			return
		}
	}
	fmt.Println("Tidak ada kartu yang cocok")
}

func sepasangKartu(a, b Domino) bool {
	return a.kiri+a.kanan+b.kiri+b.kanan == 12
}

func main() {
	data := []Domino{{1, 2}, {4, 5}, {2, 6}}
	acuan := Domino{2, 3}

	galiKartu(data, acuan)

	k1 := Domino{4, 2}
	k2 := Domino{3, 3}

	fmt.Println("Sepasang?", sepasangKartu(k1, k2))
}
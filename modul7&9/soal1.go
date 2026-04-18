package main

import "fmt"

type Titik struct {
	x, y int
}

type Lingkaran struct {
	pusat Titik
	r     int
}

func dalamLingkaran(t Titik, l Lingkaran) bool {
	dx := t.x - l.pusat.x
	dy := t.y - l.pusat.y
	return dx*dx+dy*dy <= l.r*l.r
}

func main() {
	var l1, l2 Lingkaran
	var t Titik

	fmt.Scan(&l1.pusat.x, &l1.pusat.y, &l1.r)

	fmt.Scan(&l2.pusat.x, &l2.pusat.y, &l2.r)

	fmt.Scan(&t.x, &t.y)

	in1 := dalamLingkaran(t, l1)
	in2 := dalamLingkaran(t, l2)

	if in1 && in2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if in1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if in2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
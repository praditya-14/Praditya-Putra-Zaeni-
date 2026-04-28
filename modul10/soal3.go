package main

import "fmt"

func main() {
    var n int
    fmt.Print("Masukkan banyak data berat balita: ")
    fmt.Scan(&n)

    // Gunakan slice (lebih dinamis dibanding array [100])
    berat := make([]float64, n)
    var total float64

    for i := 0; i < n; i++ {
        fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
        fmt.Scan(&berat[i])
        total += berat[i]
    }

    // Memanggil fungsi yang mengembalikan min dan max sekaligus
    min, max := cariMinMax(berat)

    fmt.Printf("\nBerat balita minimum: %.2f kg", min)
    fmt.Printf("\nBerat balita maksimum: %.2f kg", max)
    fmt.Printf("\nRerata berat balita: %.2f kg\n", total/float64(n))
}

// Fungsi simpel: satu fungsi untuk cari min dan max
func cariMinMax(data []float64) (float64, float64) {
    mi, ma := data[0], data[0]
    for _, v := range data {
        if v < mi { mi = v }
        if v > ma { ma = v }
    }
    return mi, ma
}
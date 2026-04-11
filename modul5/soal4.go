package main

import "fmt"

func pola(n int) {

    fmt.Printf("%d ", n)

    if n == 1 {
        return
    }

    pola(n - 1)

    fmt.Printf("%d ", n)
}

func main() {
    var n int
    fmt.Print("Masukkan N: ")
    fmt.Scan(&n)

    pola(n)
}
package main

import "fmt"

func main() {
	var c float64

	fmt.Print("Input suhu Celsius: ")
	fmt.Scan(&c)

	reamur := c * 4 / 5
	fahrenheit := (c * 9 / 5) + 32
	kelvin := c + 273

	fmt.Printf("Derajat Reamur: %.0f\n", reamur)
	fmt.Printf("Derajat Fahrenheit: %.0f\n", fahrenheit)
	fmt.Printf("Derajat Kelvin: %.0f\n", kelvin)
}
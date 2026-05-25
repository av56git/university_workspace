package main

import "fmt"

func main() {
	b, h := input()
	p := perimetro(b, h)
	A := area(b, h)
	fmt.Printf("Perimetro: %f\nArea: %f", p, A)
}

func input() (float64, float64) {
	var b, h float64
	fmt.Println("inserire in ordine base e altezza")
	fmt.Scan(&b, &h)
	return b, h
}

func perimetro(b float64, h float64) float64 {
	return 2 * (b + h)
}

func area(b float64, h float64) float64 {
	return b * h
}

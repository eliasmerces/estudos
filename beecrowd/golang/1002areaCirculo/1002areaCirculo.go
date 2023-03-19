package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		raio float64 = 0.00
		n    float64 = 3.14159
	)

	fmt.Scanf("%f\n", &raio)
	area := n * math.Pow(raio, 2)
	fmt.Printf("A=%.4f\n", area)
}

package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		a float64 = 0.00
		b float64 = 0.00
		c float64 = 0.00
	)

	fmt.Scanf("%f %f %f", &a, &b, &c)

	triangulo := ((a * c) / 2)
	circulo := (3.14159 * math.Pow(c, 2))
	trapezio := (((a + b) * c) / 2)
	quadrado := (b * b)
	retangulo := (a * b)

	fmt.Printf("TRIANGULO: %.3f\n", triangulo)
	fmt.Printf("CIRCULO: %.3f\n", circulo)
	fmt.Printf("TRAPEZIO: %.3f\n", trapezio)
	fmt.Printf("QUADRADO: %.3f\n", quadrado)
	fmt.Printf("RETANGULO: %.3f\n", retangulo)

}

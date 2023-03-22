package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		xUm   = 0.00
		xDois = 0.00
		yUm   = 0.00
		yDois = 0.00
	)

	fmt.Scanf("%f %f\n", &xUm, &yUm)
	fmt.Scanf("%f %f\n", &xDois, &yDois)

	pA := (xDois - xUm)
	pB := (yDois - yUm)

	pUm := math.Pow(pA, 2)
	pDois := math.Pow(pB, 2)

	resultadoA := (pUm + pDois)
	resultadoB := math.Sqrt(resultadoA)

	fmt.Printf("%.4f\n", resultadoB)

}

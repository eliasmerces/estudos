package main

import (
	"fmt"
)

func main() {
	var (
		codigoUm       int64   = 0
		codigoDois     int64   = 0
		quantidadeUm   int64   = 0
		quantidadeDois int64   = 0
		valorUm        float64 = 0.00
		valorDois      float64 = 0.00
	)

	fmt.Scanf("%d %d %f\n", &codigoUm, &quantidadeUm, &valorUm)
	fmt.Scanf("%d %d %f\n", &codigoDois, &quantidadeDois, &valorDois)

	totalUm := (float64(quantidadeUm) * valorUm)
	totalDois := (float64(quantidadeDois) * valorDois)
	totalGeral := (totalUm + totalDois)

	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", totalGeral)
}

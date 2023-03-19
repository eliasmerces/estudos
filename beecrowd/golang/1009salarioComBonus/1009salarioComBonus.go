package main

import "fmt"

func main() {
	var (
		nome          string  = ""
		salario       float64 = 0.00
		total         float64 = 0.00
		comissao      float64 = 0.00
		totalComissao float64 = 0.00
	)

	fmt.Scanf("%s\n", &nome)
	fmt.Scanf("%f\n", &salario)
	fmt.Scanf("%f\n", &total)

	comissao = (total * 0.15)
	totalComissao = (comissao + salario)

	fmt.Printf("TOTAL = R$ %.2f\n", totalComissao)
}

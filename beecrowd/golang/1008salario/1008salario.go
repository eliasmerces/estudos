package main

import (
	"fmt"
)

func main() {
	var (
		a       int64   = 0
		b       int64   = 0.00
		c       float64 = 0.00
		salario float64 = 0.00
	)

	fmt.Scanf("%d\n", &a)
	fmt.Scanf("%d\n", &b)
	fmt.Scanf("%f\n", &c)

	salario = (float64(b) * c)

	fmt.Printf("NUMBER = %d\n", a)
	fmt.Printf("SALARY = U$ %.2f\n", salario)
}

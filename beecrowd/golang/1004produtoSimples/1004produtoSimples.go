package main

import (
	"fmt"
)

func main() {
	var (
		a = 0
		b = 0
	)

	fmt.Scanf("%d\n", &a)
	fmt.Scanf("%d\n", &b)

	prod := a * b

	fmt.Printf("PROD = %d\n", prod)
}

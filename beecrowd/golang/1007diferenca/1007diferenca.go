package main

import "fmt"

func main() {
	var (
		a         int64 = 0
		b         int64 = 0
		c         int64 = 0
		d         int64 = 0
		DIFERENCA int64 = 0
	)

	fmt.Scanf("%d\n", &a)
	fmt.Scanf("%d\n", &b)
	fmt.Scanf("%d\n", &c)
	fmt.Scanf("%d\n", &d)

	DIFERENCA = ((a * b) - (c * d))
	fmt.Printf("DIFERENCA = %d\n", DIFERENCA)
}

package main

import (
	"fmt"
)

func main() {

	var (
		a int = 0
		b int = 0
	)

	fmt.Scanf("%d\n", &a)
	fmt.Scanf("%d\n", &b)

	soma := a + b

	fmt.Println("SOMA =", soma)
}

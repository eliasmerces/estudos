package main

import "fmt"

func main() {
	var (
		x int64   = 0
		y float64 = 0.00
	)

	fmt.Scanf("%d\n", &x)
	fmt.Scanf("%f\n", &y)

	media := (float64(x) / y)

	fmt.Printf("%.3f km/l\n", media)

}

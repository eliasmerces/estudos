package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		pi     float64 = 3.14159
		r      int64   = 0
		volume float64 = 0.00
	)

	fmt.Scanf("%d\n", &r)

	volume = ((4.00 / 3.00) * pi * math.Pow(float64(r), 3))

	fmt.Printf("VOLUME = %.3f\n", volume)
}

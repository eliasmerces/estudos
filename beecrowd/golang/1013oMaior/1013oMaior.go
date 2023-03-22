package main

import (
	"fmt"
	"math"
)

func main() {
	var (
		a     float64 = 0.00
		b     float64 = 0.00
		c     float64 = 0.00
		absAB float64 = 0.00
		absC  float64 = 0.00
	)

	fmt.Scanf("%f %f %f\n", &a, &b, &c)

	absAB = (a - b)
	maiorAB := ((a + b + math.Abs(absAB)) / 2)
	absC = (maiorAB - c)
	maiorC := ((maiorAB + c + math.Abs(absC)) / 2)

	fmt.Printf("%.0f eh o maior\n", maiorC)
}

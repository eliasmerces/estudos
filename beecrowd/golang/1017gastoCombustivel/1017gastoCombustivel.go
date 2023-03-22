package main

import "fmt"

func main() {
	var (
		horas     int64 = 0
		distancia int64 = 0
	)

	fmt.Scanf("%d\n", &horas)
	fmt.Scanf("%d\n", &distancia)

	gastoCombustivel := (float64(distancia) / 12.00)
	total := (float64(horas) * gastoCombustivel)

	fmt.Printf("%.3f\n", total)
}

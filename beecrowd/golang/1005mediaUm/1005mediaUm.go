package main

import "fmt"

func main() {
	var (
		a     float64 = 0.00
		b     float64 = 0.00
		media float64 = 0.00
	)

	fmt.Scanf("%f\n", &a)
	fmt.Scanf("%f\n", &b)

	if a <= 10.00 {
		if a >= 0.00 {
			if b <= 10.00 {
				if b >= 0.00 {
					media = ((a * 3.5) + (b * 7.5)) / 11
					fmt.Printf("MEDIA = %.5f\n", media)
				} else {
					fmt.Println("Nota informada invalida")
				}
			} else {
				fmt.Println("Nota informada invalida")
			}
		} else {
			fmt.Println("Nota informada invalida")
		}
	} else {
		fmt.Println("Nota informada invalida")
	}
}

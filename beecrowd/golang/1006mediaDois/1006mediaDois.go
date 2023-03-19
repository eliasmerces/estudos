package main

import "fmt"

func main() {
	var (
		a     float64 = 0.00
		b     float64 = 0.00
		c     float64 = 0.00
		media float64 = 0.00
	)

	fmt.Scanf("%f\n", &a)
	fmt.Scanf("%f\n", &b)
	fmt.Scanf("%f\n", &c)

	if a <= 10 {
		if a >= 0 {
			if b <= 10 {
				if b >= 0 {
					if c <= 10 {
						if c >= 0 {
							media = (((a * 2) + (b * 3) + (c * 5)) / 10)
							fmt.Printf("MEDIA = %.1f\n", media)
						} else {
							fmt.Println("Nota invalida")
						}
					} else {
						fmt.Println("Nota invalida")
					}
				} else {
					fmt.Println("Nota invalida")
				}
			} else {
				fmt.Println("Nota invalida")
			}
		} else {
			fmt.Println("Nota invalida")
		}
	} else {
		fmt.Println("Nota invalida")
	}
}

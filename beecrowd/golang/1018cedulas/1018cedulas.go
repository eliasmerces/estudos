package main

import "fmt"

func main() {
	var (
		dinh  int64 = 0
		valor int64 = 0
		cem   int64 = 0
		ciq   int64 = 0
		vin   int64 = 0
		dez   int64 = 0
		cin   int64 = 0
		doi   int64 = 0
	)

	fmt.Scanf("%d\n", &dinh)

	if dinh > 0 {
		if dinh < 1000000 {

			valor = dinh

			cem = (valor / 100)
			valor = (valor - (cem * 100))
			ciq = (valor / 50)
			valor = (valor - (ciq * 50))
			vin = (valor / 20)
			valor = (valor - (vin * 20))
			dez = (valor / 10)
			valor = (valor - (dez * 10))
			cin = (valor / 5)
			valor = (valor - (cin * 5))
			doi = (valor / 2)
			valor = (valor - (doi * 2))

			fmt.Printf("%d\n", dinh)
			fmt.Printf("%d nota(s) de R$ 100,00\n", cem)
			fmt.Printf("%d nota(s) de R$ 50,00\n", ciq)
			fmt.Printf("%d nota(s) de R$ 20,00\n", vin)
			fmt.Printf("%d nota(s) de R$ 10,00\n", dez)
			fmt.Printf("%d nota(s) de R$ 5,00\n", cin)
			fmt.Printf("%d nota(s) de R$ 2,00\n", doi)
			fmt.Printf("%d nota(s) de R$ 1,00\n", valor)

		} else {
			fmt.Println("Valor informado invalido!")
		}
	} else {
		fmt.Println("Valor informado invalido!")
	}

}

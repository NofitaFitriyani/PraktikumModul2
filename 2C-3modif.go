package main

import "fmt"

func main() {
	var bilfaktor int
	fmt.Print("Bilangan: ")
	fmt.Scan(&bilfaktor)

	isPrima := true

	fmt.Print("Faktor : ")
	for i := 1; i <= bilfaktor; i++ {
		if bilfaktor%i == 0 {
			fmt.Print(i, " ")
			if i > 1 && i < bilfaktor {
				isPrima = false
			}
		}
	}

	fmt.Println("\nPrima: ", isPrima)
}

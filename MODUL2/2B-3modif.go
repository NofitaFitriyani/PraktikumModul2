package main

import "fmt"

func main() {
	var beratKantongA, beratKantongB float64
	var isGoyang bool

	for {
		fmt.Print("Masukkan berat belanjaan kedua kantong: ")
		fmt.Scan(&beratKantongA, &beratKantongB)

		difference := beratKantongA - beratKantongB
		if difference < 0 {
			difference = -difference
		}
		isGoyang = difference >= 9

		if beratKantongA+beratKantongB > 150 {
			fmt.Print("Proses selesai")
			break
		}

		fmt.Println("Sepeda motor pak Andi akan oleng:", isGoyang)
	}
}

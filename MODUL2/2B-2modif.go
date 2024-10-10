package main

import "fmt"

func main() {
	var jumlahBunga int
	var namaBunga, rangkaianBunga string

	for {
		fmt.Print("Bunga ke-", jumlahBunga+1, ": ")
		fmt.Scan(&namaBunga)

		if namaBunga == "SELESAI" {
			break
		}

		rangkaianBunga += namaBunga + " - "
		jumlahBunga++
	}

	fmt.Print("Rangkaian: ", rangkaianBunga)
	fmt.Print("Total Bunga: ", jumlahBunga)
}

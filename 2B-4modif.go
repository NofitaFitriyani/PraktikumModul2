package main

import (
	"fmt"
	"math"
)

func main() {
	var k int
	var fungsi float64 = 1.0
	fmt.Print("Nilai K = ")
	fmt.Scan(&k)
	for i := 0; i <= k; i++ {
		pembilang := math.Pow(4*float64(i)+2, 2)
		penyebut := ((4*float64(i) + 1) * (4*float64(i) + 3))
		fungsi *= pembilang / penyebut
	}
	fmt.Printf("nilai f(k) = %.10f", fungsi)
}

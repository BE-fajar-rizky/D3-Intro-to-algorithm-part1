package main

import "fmt"

func main() {
	/// cara pertama ya mas
	var alas float64 = 20
	var tinggi float64 = 25

	var luas float64 = alas * tinggi / 2

	fmt.Println("luas segitiga adalah = ", luas)

	///cara kedua

	Alas := 20
	Tinggi := 25
	var Luas float64
	Luas = float64(Alas) * float64(Tinggi) / 2

	fmt.Println("luas segitiganya adalah = ", Luas)
}

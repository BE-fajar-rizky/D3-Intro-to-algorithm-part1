package main

import "fmt"

func main() {
	/// cara pertama ya mas
	var Pi float64 = 3.14
	var r float64 = 4
	var T float64 = 20

	var Lp float64 = 2 * Pi * r * (r + T)

	fmt.Println("luas permukaan tabung adalah = ", Lp)

	///cara kedua

	pi := 3.14
	R := 4
	t := 20

	var lp float64

	lp = 2 * float64(pi) * float64(R) * (float64(R) + float64(t))

	fmt.Println("luas Permukaan tabungnya adalah = ", lp)

}

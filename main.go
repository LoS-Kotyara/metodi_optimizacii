package main

import (
	"fmt"
	"metodi_optimizacii/methods"
)

func main() {
	regularSimplex, N := methods.RegularSimplex(methods.EPS, methods.BasePoint)

	fmt.Printf("Simplex count:\t%v\n", N)
	fmt.Println("Result:")
	regularSimplex.Print()
}

package main

import (
	"fmt"
	"metodi_optimizacii/methods"
)

func main() {
	//n := 3 // E_2

	l := 0.001
	//l := 2.0
	x0 := [2]float64{-2, 1}
	//x0 := [2]float64{0, 0}

	regularSimplex := methods.RegularSimplex(l, x0)

	fmt.Printf("Result:\t%v\n", regularSimplex)
}

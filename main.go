package main

import (
	"fmt"
	"metodi_optimizacii/methods"
)

func main() {
	sepLength := 40

	methods.Separator(sepLength)
	fmt.Println("Exact values:")
	fmt.Println("\tx1 = ", methods.ExactX1)
	fmt.Println("\tx2 = ", methods.ExactX2)
	fmt.Println("\tval = ", methods.ExactMin)

	methods.Separator(sepLength)
	regularSimplex, N := methods.RegularSimplex(methods.EPS, methods.BasePoint)
	fmt.Println("Regular:")
	fmt.Printf("Simplex count:\t%v\n", N)
	fmt.Println("Result:")
	regularSimplex.Print()

	methods.Separator(sepLength)
	nelderMead, N := methods.NelderMead(methods.EPS, methods.BasePoint)
	fmt.Println("NelderMead:")
	fmt.Printf("Simplex count:\t%v\n", N)
	fmt.Println("Result:")
	nelderMead.Print()

	methods.Separator(sepLength)

	val, point, iter := methods.CoordinateDescent()
	fmt.Println("CoordinateDescent:")
	fmt.Println("Result:")
	point.Print()
	fmt.Println("\tval = ", val)
	fmt.Println("\titerations: ", iter)

	methods.Separator(sepLength)
	val, point, iter = methods.HookeJeeves()
	fmt.Println("HookeJeeves:")
	fmt.Println("Result:")
	point.Print()
	fmt.Println("\tval = ", val)
	fmt.Println("\titerations: ", iter)

	methods.Separator(sepLength)
	val, point, iter = methods.RandomSearch()
	fmt.Println("RandomSearch:")
	fmt.Println("Result:")
	point.Print()
	fmt.Println("\tval = ", val)
	fmt.Println("\titerations: ", iter)

	methods.Separator(sepLength)
}

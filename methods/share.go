package methods

import "math"

//const EPS = 0.001

func f(x1, x2 float64) float64 {
	//return 43*x1*x1 + 84*x1*x2 + 43*x2*x2 + 15*x1 - 26*x2 + 42
	return 6*x1*x1 - 4*x1*x2 + 3*x2*x2 + 4*math.Sqrt(5)*(x1+2*x2) + 22
}

type Point [2]float64

type SimplexVertex struct {
	fVal  float64
	point Point
}

type Simplex struct {
	vertexes [3]SimplexVertex
	pointer  int
}

func (simplex Simplex) Sort() Simplex {
	for i := 0; i < len(simplex.vertexes); i++ {
		for j := 0; j < len(simplex.vertexes)-i-1; j++ {
			if simplex.vertexes[j].fVal > simplex.vertexes[j+1].fVal {
				simplex.vertexes[j], simplex.vertexes[j+1] = simplex.vertexes[j+1], simplex.vertexes[j]
			}
		}
	}

	return simplex
}

type SimplexHist []Simplex

package methods

import "fmt"

const (
	EPS = 0.001
)

var (
	BasePoint = Point{-2, 1}
)

func f(point Point) float64 {
	x1 := point[0]
	x2 := point[1]
	return 43*x1*x1 + 84*x1*x2 + 43*x2*x2 + 15*x1 - 26*x2 + 42
}

type (
	Point         [2]float64
	SimplexHist   []Simplex
	SimplexVertex struct {
		fVal  float64
		point Point
	}
	Simplex struct {
		vertexes [3]SimplexVertex
		pointer  int
	}
)

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

func (v SimplexVertex) Print() {
	for i := 0; i < len(v.point); i++ {
		fmt.Printf("\tx%v = %v\n", i, v.point[i])
	}
	fmt.Println("\tval = ", v.fVal)
}

func (simplex Simplex) Print() {
	for _, v := range simplex.vertexes {
		v.Print()
	}
}

func (hist SimplexHist) Print() {
	for i, v := range hist {
		fmt.Printf("%d:\n", i)
		printSimplex(v)
		fmt.Println("\t", v.pointer)
	}
}

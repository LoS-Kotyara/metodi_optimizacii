package methods

import (
	"fmt"
	"math"
	"strings"
)

const (
	EPS      = 0.001
	N        = 2
	ExactX1  = -10.217647058823529
	ExactX2  = 10.282352941176470
	ExactMin = -168.30294117647058
)

var (
	BasePoint = Point{-2, 1}
)

// min{43 x_1^2 + 84 x_1 x_2 + 43 x_2^2 + 15 x_1 - 26 x_2 + 42}≈-168.30294117647058
//	at (x_1, x_2)≈(-10.217647058823529, 10.282352941176470)
func f(point Point) float64 {
	x1 := point[0]
	x2 := point[1]

	return 43*x1*x1 + 84*x1*x2 + 43*x2*x2 + 15*x1 - 26*x2 + 42
	//return 6*x1*x1 - 4*x1*x2 + 3*x2*x2 + 4*math.Sqrt(5)*(x1+2*x2) + 22
}

type (
	Point         [2]float64
	SimplexHist   []Simplex
	SimplexVertex struct {
		fVal  float64
		point Point
	}
	Simplex struct {
		vertexes [N + 1]SimplexVertex
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

func (p Point) Print() {
	for i := 0; i < len(p); i++ {
		fmt.Printf("\tx%v = %v\n", i, p[i])
	}
}

func (v SimplexVertex) Print() {
	v.point.Print()
	fmt.Println("\tval = ", v.fVal)
}

func (simplex Simplex) Print() {
	for i, v := range simplex.vertexes {
		fmt.Println("\t", i, ":")
		v.Print()
	}
}

func (hist SimplexHist) Print() {
	for i, v := range hist {
		fmt.Printf("%d:\n", i+1)
		v.Print()
		fmt.Println("\t", v.pointer)
	}
}

func newBaseSimplex(x0 Point) Simplex {
	l := EPS
	n := N
	simplex := new(Simplex)
	for _i := 0; _i < n+1; _i++ {
		i := float64(_i) + 1

		var point Point
		for _j, x := range x0 {
			j := float64(_j) + 1.0

			if j < i-1 {
				point[_j] = x
			} else if j == i-1 {
				point[_j] =
					x + math.Sqrt(j/(2*(j+1)))*l
			} else if j > i-1 {
				point[_j] =
					x - 1/(math.Sqrt(2*j*(j+1)))*l

			}
		}
		simplex.vertexes[_i] = SimplexVertex{fVal: f(point), point: point}
	}
	simplex.pointer = n // greatest func value index
	return simplex.Sort()
}

func _reflectSimplexVertex(simplex *Simplex) Simplex {
	n := N
	pointer := simplex.pointer

	toReflect := &simplex.vertexes[pointer]
	var length [2]float64
	for i := 0; i < n; i++ {
		if i != pointer {
			length[0] += simplex.vertexes[i].point[0] - toReflect.point[0]
			length[1] += simplex.vertexes[i].point[1] - toReflect.point[1]
		}
	}

	newPoint := Point{
		toReflect.point[0] + length[0],
		toReflect.point[1] + length[1],
	}

	newVertex := SimplexVertex{point: newPoint, fVal: f(newPoint)}
	var newSimplex Simplex
	j := 0
	for i := 0; i < n; i++ {
		if i != pointer {
			newSimplex.vertexes[j] = simplex.vertexes[i]
			j++
		}
	}
	newSimplex.vertexes[j] = newVertex

	return newSimplex
}

func Separator(len int) {
	fmt.Println(strings.Repeat("*", len))
}

func (simplex Simplex) reflectVertex() Simplex {
	n := N
	point := Point{0, 0}

	pointer := simplex.pointer

	for i := 0; i < n+1; i++ {
		if i != pointer {
			point[0] += simplex.vertexes[i].point[0]
			point[1] += simplex.vertexes[i].point[1]
		}
	}

	point[0] *= 2 / float64(n)
	point[1] *= 2 / float64(n)

	point[0] -= simplex.vertexes[pointer].point[0]
	point[1] -= simplex.vertexes[pointer].point[1]

	newVertex := SimplexVertex{
		fVal:  f(point),
		point: point,
	}

	var newSimplex Simplex
	j := 0
	for i := 0; i < n+1; i++ {
		if i != pointer {
			newSimplex.vertexes[j] = simplex.vertexes[i]
			j++
		}
	}
	newSimplex.vertexes[j] = newVertex

	return newSimplex.Sort()
}

func reflectSimplexVertex(simplex *Simplex) Simplex {
	n := N
	point := Point{0, 0}

	pointer := simplex.pointer

	for i := 0; i < n+1; i++ {
		if i != pointer {
			point[0] += simplex.vertexes[i].point[0]
			point[1] += simplex.vertexes[i].point[1]
		}
	}

	point[0] *= 2 / float64(n)
	point[1] *= 2 / float64(n)

	point[0] -= simplex.vertexes[pointer].point[0]
	point[1] -= simplex.vertexes[pointer].point[1]

	newVertex := SimplexVertex{
		fVal:  f(point),
		point: point,
	}

	var newSimplex Simplex
	j := 0
	for i := 0; i < n+1; i++ {
		if i != pointer {
			newSimplex.vertexes[j] = simplex.vertexes[i]
			j++
		}
	}
	newSimplex.vertexes[j] = newVertex

	return newSimplex
}

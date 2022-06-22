package methods

import (
	"math"
)

func reflectSimplexVertex(simplex *Simplex, n int) Simplex {
	//n := 3 // E_2
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

	//pointer := simplex.pointer
	//
	//toReflect := &simplex.vertexes[pointer]
	//var length [2]float64
	//for i := 0; i < n; i++ {
	//	if i != pointer {
	//		length[0] += simplex.vertexes[i].point[0] - toReflect.point[0]
	//		length[1] += simplex.vertexes[i].point[1] - toReflect.point[1]
	//	}
	//}
	//
	//newPoint := Point{
	//	toReflect.point[0] + length[0],
	//	toReflect.point[1] + length[1],
	//}
	//
	//newVertex := SimplexVertex{point: newPoint, fVal: f(newPoint)}
	//var newSimplex Simplex
	//j := 0
	//for i := 0; i < n; i++ {
	//	if i != pointer {
	//		newSimplex.vertexes[j] = simplex.vertexes[i]
	//		j++
	//	}
	//}
	//newSimplex.vertexes[j] = newVertex

	return newSimplex
}

func RegularSimplex(l float64, x0 [2]float64) (SimplexVertex, int) {
	n := N

	var simplexHist SimplexHist

	var simplex Simplex
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
			} else if j > i-1 { // j > i - 1
				point[_j] =
					x - 1/(math.Sqrt(2*j*(j+1)))*l

			}
		}
		simplex.vertexes[_i] = SimplexVertex{fVal: f(point), point: point}
	}
	simplex.pointer = n // greatest func value index
	simplexHist = append(simplexHist, simplex.Sort())

	for true {
		lastSimplex := &simplexHist[len(simplexHist)-1]

		newSimplex := reflectSimplexVertex(lastSimplex, n).Sort()

		if newSimplex.vertexes[2].fVal < lastSimplex.vertexes[lastSimplex.pointer].fVal {
			newSimplex.pointer = n
			simplexHist = append(simplexHist, newSimplex)
			continue
		} else {
			if lastSimplex.pointer > 0 {
				lastSimplex.pointer = lastSimplex.pointer - 1
				continue
			}
			break
		}
	}

	//simplexHist.Print()
	return simplexHist[len(simplexHist)-1].vertexes[0], len(simplexHist)
}

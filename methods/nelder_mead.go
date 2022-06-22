package methods

const (
	alpha = 0.5
	beta  = 1.0
	gamma = 2.0
)

func minFromArray(a ...float64) float64 {
	min := a[0]
	for _, v := range a {
		if v < min {
			min = v
		}
	}

	return min
}

func nextSimplex(simplex Simplex) Simplex {
	n := N
	sumOfPoints := Point{0, 0}

	pointer := simplex.pointer

	for i := 0; i < n+1; i++ {
		if i != pointer {
			sumOfPoints[0] += simplex.vertexes[i].point[0]
			sumOfPoints[1] += simplex.vertexes[i].point[1]
		}
	}

	// contraction (inner)
	z1 := Point{(1-alpha)/float64(n)*sumOfPoints[0] + alpha*simplex.vertexes[pointer].point[0],
		(1-alpha)/float64(n)*sumOfPoints[1] + alpha*simplex.vertexes[pointer].point[1]}

	// contraction (outer)
	z2 := Point{(1+alpha)/float64(n)*sumOfPoints[0] - alpha*simplex.vertexes[pointer].point[0],
		(1+alpha)/float64(n)*sumOfPoints[1] - alpha*simplex.vertexes[pointer].point[1]}

	// reflection
	z3 := Point{(1+beta)/float64(n)*sumOfPoints[0] - beta*simplex.vertexes[pointer].point[0],
		(1+beta)/float64(n)*sumOfPoints[1] - beta*simplex.vertexes[pointer].point[1]}

	// expansion
	z4 := Point{(1+gamma)/float64(n)*sumOfPoints[0] - gamma*simplex.vertexes[pointer].point[0],
		(1+gamma)/float64(n)*sumOfPoints[1] - gamma*simplex.vertexes[pointer].point[1]}

	fZ1, fZ2, fZ3, fZ4 := f(z1), f(z2), f(z3), f(z4)
	min := minFromArray(fZ1, fZ2, fZ3, fZ4)

	var newVertex SimplexVertex
	switch min {
	case fZ1:
		newVertex.fVal = fZ1
		newVertex.point = z1
	case fZ2:
		newVertex.fVal = fZ2
		newVertex.point = z2
	case fZ3:
		newVertex.fVal = fZ3
		newVertex.point = z3
	case fZ4:
		newVertex.fVal = fZ4
		newVertex.point = z4
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

func NelderMead(l float64, x0 Point) (SimplexVertex, int) {
	n := N
	simplexHist := make(SimplexHist, 0)

	s1 := newBaseSimplex(x0)

	simplexHist = append(simplexHist, s1)

	for true {
		lastSimplex := &simplexHist[len(simplexHist)-1]

		newSimplex := nextSimplex(*lastSimplex)

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

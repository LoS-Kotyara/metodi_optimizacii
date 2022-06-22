package methods

func RegularSimplex(l float64, x0 Point) (SimplexVertex, int) {
	n := N
	simplexHist := make(SimplexHist, 0)

	s1 := newBaseSimplex(x0)

	simplexHist = append(simplexHist, s1)

	for true {
		lastSimplex := &simplexHist[len(simplexHist)-1]
		newSimplex := lastSimplex.reflectVertex()

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

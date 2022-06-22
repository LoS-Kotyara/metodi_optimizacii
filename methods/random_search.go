package methods

import (
	"math/rand"
	"time"
)

func RandomSearch() (float64, Point, int) {
	s1 := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(s1)

	alpha := 1.0
	gamma := 2.0
	M := 30 * N
	x := Point{randGen.NormFloat64(), randGen.NormFloat64()}
	fX := f(x)

	j := 1
	iter := 0
	for true {
		iter++
		y := Point{}

		for i := 0; i < N; i++ {
			xi := randGen.NormFloat64()
			y[i] = x[i] + alpha*xi
		}

		fY := f(y)

		if fY < fX {
			x, fX = y, fY
		} else {
			j++

			if j > M {
				if alpha < EPS {
					break
				} else {
					alpha /= gamma
					j = 1
				}
			}
		}
	}

	return fX, x, iter
}

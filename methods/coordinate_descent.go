package methods

import (
	"math"
)

func bitwiseSearch(x Point, ord int) Point {
	a := x[ord] - 10
	b := x[ord] + 10

	d := (b - a) / 4.0

	x1 := x
	x2 := x
	x1[ord] = a

	for math.Abs(d) > EPS {
		x2[ord] = x1[ord] + d
		if f(x1) < f(x2) {
			d = -(d / 4)
			x1 = x2
		} else {
			x1 = x2
		}
	}

	return x1
}

func CoordinateDescent() (float64, Point, int) {
	x := BasePoint
	var xp Point

	const MAX_ITER = 500
	for iter := 0; iter < MAX_ITER; iter++ {
		xp = x

		for i := 0; i < N; i++ {
			x = bitwiseSearch(x, i)
		}

		if math.Abs(f(x)-f(xp)) < EPS {
			return f(x), x, iter
		}
	}
	return f(x), x, MAX_ITER
}

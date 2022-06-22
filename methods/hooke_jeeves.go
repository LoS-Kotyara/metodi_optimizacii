package methods

func hookeJeevesSearch(x Point, delta float64) Point {
	n := N
	xO := x

	for j := 0; j < n; j++ {
		y := xO
		y[j] = xO[j] - delta
		//fmt.Println("1 level", j, y)

		if f(xO) <= f(y) {
			//fmt.Println("2 level", j, y)
			y[j] = xO[j] + delta

			if f(xO) <= f(y) {
				//fmt.Println("3 level", j, xO)
				continue
			} else {
				xO = y
			}
		} else {
			xO = y
		}
	}

	return xO
}

func HookeJeeves() (fVal float64, point Point, iter int) {
	xb := point
	xp := point
	var xn Point

	delta := 0.5
	gamma := 2.0
	iter = 0
	for true {
		iter++
		if delta < EPS {
			break
		}

		xn = hookeJeevesSearch(xp, delta)

		if fN, fB := f(xn), f(xb); fN < fB {
			for i := 0; i < N; i++ {
				xp[i] = 2*xn[i] - xb[i]
				xb[i] = xn[i]
			}
		} else {
			delta /= gamma
			xp = xb
		}

	}
	return f(xb), xb, iter
}

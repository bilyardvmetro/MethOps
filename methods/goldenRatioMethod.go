package methods

import "fmt"

var stepGR = 1

const GR1 = 0.382
const GR2 = 0.618

func SolveByGoldenRatio(f func(float64) float64, a float64, b float64, eps float64) (float64, float64) {
	x1 := a + GR1*(b-a)
	x2 := a + GR2*(b-a)
	y1 := f(x1)
	y2 := f(x2)

	for b-a > eps {
		fmt.Printf("Шаг: %d\nОтрезок: [%f; %f]\n", stepGR, a, b)

		//fmt.Printf("x1=%f, x2=%f, y1=%f, y2=%f\n", x1, x2, y1, y2)

		if y1 < y2 {
			b = x2
			x2 = x1
			y2 = y1
			x1 = a + GR1*(b-a)
			y1 = f(x1)

			//fmt.Printf("y1 < y2. b = %f\n", x2)
		} else {
			a = x1
			x1 = x2
			y1 = y2
			x2 = a + GR2*(b-a)
			y2 = f(x2)

			//fmt.Printf("y1 >= y2. a = %f\n", x1)
		}
		fmt.Printf("a = %f, b = %f\n", a, b)
		//fmt.Printf("b - a = %f\n", b-a)
		stepGR++
	}

	fmt.Printf("\nb - a < ε. Минимум с погрешностью ε = %f лежит на середине этого отрезка\n", eps)
	xM := (a + b) / 2
	yM := f(xM)
	fmt.Printf("Минимум в точке x_m = %f. Значение функции f(x_m) = y_m = %f\n", xM, yM)

	return xM, yM
}

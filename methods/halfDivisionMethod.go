package methods

import "fmt"

var stepHD = 1

func SolveByHalfDivision(f func(float64) float64, a float64, b float64, eps float64) (float64, float64) {
	for b-a > 2*eps {
		fmt.Printf("Шаг: %d\nОтрезок: [%f; %f]\n", stepHD, a, b)

		x1 := (a + b - eps) / 2
		x2 := (a + b + eps) / 2
		y1 := f(x1)
		y2 := f(x2)

		//fmt.Printf("x1=%f, x2=%f, y1=%f, y2=%f\n", x1, x2, y1, y2)

		if y1 > y2 {
			//fmt.Printf("y1 > y2. Новый отрезок: [a; %f]\n", x1)
			a = x1
		} else {
			//fmt.Printf("y1 <= y2. Новый отрезок: [%f; b]\n", x2)
			b = x2
		}
		fmt.Printf("a = %f, b = %f\n", a, b)
		//fmt.Printf("b - a = %f\n", b-a)
		stepHD++
	}

	fmt.Printf("\nb - a < 2ε. Минимум с погрешностью ε = %f лежит на середине этого отрезка\n", eps)
	xM := (a + b) / 2
	yM := f(xM)
	fmt.Printf("Минимум в точке x_m = %f. Значение функции f(x_m) = y_m = %f\n", xM, yM)

	return xM, yM
}

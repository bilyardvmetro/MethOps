package methods

import (
	"fmt"
	"math"
)

var stepN = 1

func SolveByNewtonMethod(f func(float64) float64, a float64, b float64, eps float64) (float64, float64) {
	fFd := FirstDerivative()
	fSd := SecondDerivative()

	x := (a + b) / 2
	fmt.Printf("Начинаем с середины отрезка x0 = %f\n", x)

	for math.Abs(fFd(x)) > eps {
		fmt.Printf("Шаг: %d\n", stepN)

		x = x - (fFd(x) / fSd(x))

		fmt.Printf("x = %f, f'(x) = %f\n", x, fFd(x))
		stepN++
	}

	fmt.Printf("\n|f'(x)| <= ε. Найден минимум с погрешностью ε = %f\n", eps)
	xM := x
	yM := f(x)
	fmt.Printf("Минимум в точке x_m = %f. Значение функции f(x_m) = y_m = %f\n", xM, yM)

	return xM, yM
}

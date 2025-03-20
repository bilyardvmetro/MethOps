package methods

import (
	"fmt"
	"math"
)

var stepC = 1

func SolveByChordMethod(f func(float64) float64, a float64, b float64, eps float64) (float64, float64) {
	fFd := FirstDerivative()
	fFdB := fFd(b)
	fFdA := fFd(a)
	var xWave float64

	for math.Abs(fFd(xWave)) > eps {
		fmt.Printf("Шаг: %d\nОтрезок: [%f; %f]\n", stepC, a, b)

		xWave = a - (fFdA/(fFdA-fFdB))*(a-b)
		fXWave := fFd(xWave)

		fmt.Printf("x с волной = %f, f(x c волной) = %f\n", xWave, fXWave)

		if fXWave > 0 {
			b = xWave
			fFdB = fXWave
			fmt.Printf("f(x c волной) > 0. Значит b = %f. f'(b) = f'(x с волной) = %f\n", xWave, fXWave)
		} else {
			a = xWave
			fFdA = fXWave
			fmt.Printf("f(x c волной) <= 0. Значит a = %f. f'(a) = f'(x с волной) = %f\n", xWave, fXWave)
		}

		fmt.Printf("a = %f, b = %f\n", a, b)
		stepC++
	}

	fmt.Printf("\n|f'(x)| <= ε. Найден минимум с погрешностью ε = %f\n", eps)
	xM := xWave
	yM := f(xWave)
	fmt.Printf("Минимум в точке x_m = %f. Значение функции f(x_m) = y_m = %f\n", xM, yM)

	return xM, yM
}

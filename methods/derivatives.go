package methods

import "math"

func FirstDerivative() func(float64) float64 {
	return func(x float64) float64 {
		return 3*(math.Pow(x, 2)) - 3*math.Cos(x)
	}
}

func SecondDerivative() func(float64) float64 {
	return func(x float64) float64 {
		return 6*x + 3*math.Sin(x)
	}
}

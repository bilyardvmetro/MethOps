package main

import (
	"OptMethLab23/methods"
	"fmt"
	"math"
)

func f(x float64) float64 {
	return math.Pow(x, 3) - 3*math.Sin(x)
}

type methodFunction func(f func(float64) float64, a, b, eps float64) (float64, float64)

func main() {
	a, b, eps := 0.0, 1.0, 0.0001
	m := map[string]methodFunction{
		"Вычисление методом половинного деления": methods.SolveByHalfDivision,
		"Вычисление методом золотого сечения":    methods.SolveByGoldenRatio,
		"Вычисление методом хорд":                methods.SolveByChordMethod,
		"Вычисление методом Ньютона":             methods.SolveByNewtonMethod,
	}

	for k := range m {
		fmt.Println(k + ":")
		xM, yM := m[k](f, a, b, eps)
		fmt.Printf("Точка минимума: %f; Минимум: %f\n", xM, yM)
		fmt.Println("========================================")
	}

	fmt.Println("Вычисление методом квадратной аппроксимации: ")
	xM, yM := methods.QuadraticApproximation(f, a, (a+b)/2, b, 0.0001)
	fmt.Printf("Точка минимума: %f; Минимум: %f\n", xM, yM)
	fmt.Println("========================================")
}

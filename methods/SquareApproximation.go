package methods

import (
	"fmt"
	"math"
)

var stepQA = 1

// QuadraticApproximation реализует метод квадратичной аппроксимации для поиска минимума функции
func QuadraticApproximation(f func(float64) float64, x1, x2, x3, eps float64) (float64, float64) {
	for math.Abs(x3-x1) > eps {
		fmt.Printf("Шаг: %d\nОтрезок: [%f; %f]\nТекущий минимум: %f\n", stepQA, x1, x3, x2)
		f1, f2, f3 := f(x1), f(x2), f(x3)

		// Вычисляем вершину параболы, аппроксимирующей функцию
		numerator := (math.Pow(x2, 2)-math.Pow(x3, 2))*f1 + (math.Pow(x3, 2)-math.Pow(x1, 2))*f2 + (math.Pow(x1, 2)-math.Pow(x2, 2))*f3
		denominator := (x2-x3)*f1 + (x3-x1)*f2 + (x1-x2)*f3
		if denominator == 0 {
			// Избегаем деления на ноль
			x1 = min(x1, x2, x3)
			continue
		}
		xMin := 0.5 * numerator / denominator

		// Обновляем точки для следующей итерации
		if xMin > x2 {
			fmt.Printf("x_min = %f > x2 = %f\nСледовательно x1 = x2 = %f и x2 = x_min = %f\n", xMin, x2, x2, xMin)
			x1, x2 = x2, xMin
		} else {
			fmt.Printf("x_min = %f <= x2 = %f\nСледовательно x3 = x2 = %f и x2 = x_min = %f\n", xMin, x2, x2, xMin)
			x2, x3 = xMin, x2
		}
		stepQA++
		fmt.Printf("x3 - x1 = %f\n", x3-x1)
	}
	return x2, f(x2)
}

//func SolveBySquareApproximation(f func(float64) float64, yEps float64, xEps float64) (float64, float64) {
//	x1 := 0.0
//	deltaX := 0.5
//
//	for {
//		x2 := x1 + deltaX
//		y1 := f(x1)
//		y2 := f(x2)
//
//		x3 := 0.0
//		if y1 > y2 {
//			x3 = x1 + 2*deltaX
//		} else {
//			x3 = x1 - deltaX
//		}
//
//		y3 := f(x3)
//
//		for {
//			fMin := min(y1, y2, y3)
//			xMin := min(x1, x2, x3)
//
//			if (x2-x3)*y1+(x3-x1)*y2+(x1-x2)*y3 == 0 {
//				x1 = xMin
//				break
//			}
//			xStripe := 0.5 * (((math.Pow(x2, 2)-math.Pow(x3, 2))*y1 + (math.Pow(x3, 2)-math.Pow(x1, 2))*y2 + (math.Pow(x1, 2)-math.Pow(x2, 2))*y3) / ((x2-x3)*y1 + (x3-x1)*y2 + (x1-x2)*y3))
//			fXStripe := f(xStripe)
//
//			yCondition := math.Abs((fMin-fXStripe)/fXStripe) < yEps
//			xCondition := math.Abs((xMin-xStripe)/xStripe) < xEps
//
//			if xCondition && yCondition {
//				// а
//				xM := xMin
//				return xM, f(xM)
//			} else if xCondition || yCondition {
//				if xStripe >= x1 && xStripe <= x3 {
//					// б
//					x2 = min(xStripe, xMin)
//					x1 = min(x1, x2, x3)
//					x3 = max(x1, x2, x3)
//					continue
//				} else {
//					// в
//					x1 = xStripe
//					break
//				}
//			}
//		}
//	}
//}

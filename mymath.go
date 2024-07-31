package mymath

import "math"

// Sqrt - функция для вычисления квадратного корня из числа
func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}

func Ceil(x float64) float64 {
	return math.Ceil(x)
}

func Abs(x float64) float64{
	return math.Abs(x)
}

func Yn(n int, x float64) float64{
	return math.Yn(n,x)
}

func Floor(x float64) float64 {
	return math.Floor(x)
}

func Pow(x, y float64) float64 {
	return math.Pow(x, y)
}

func Max(x, y float64) float64 {
	return math.Max(x, y)
}

func Acos(x float64) float64 {
	return math.Acos(x)
}
func Acosh(x float64) float64 {
	return math.Acosh(x)
}
func Asin(x float64) float64 {
	return math.Asin(x)
}

func Min(x, y float64) float64 {
	return math.Min(x, y)
}

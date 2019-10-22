package main

import (
	"fmt"
	"math"
	//	"runtime"
	//"time"
)

func needInt(x int) int           { return x*100 + 1 }
func needFloat(x float64) float64 { return x * 0.1 }

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		defer fmt.Println(
			z,
			i,
		)
		z = z - (z*z-x)/(2*z)
	}
	return z
}

func main() {
	defer fmt.Println("defer..." , Sqrt(2))
	fmt.Println("hello")
}

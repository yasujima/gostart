package main

import (
	"fmt"
	"math"
)

type Ver struct {
	X, Y float64
}

func (v *Ver) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v* Ver) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func BB(v Ver)  {
	v.X = v.X * 100
	v.Y = v.Y * 100
}

func main() {
	v := Ver{3,4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(&v))

	p := &Ver{4,3}
	fmt.Println((p).Abs())
	fmt.Println(AbsFunc(p))
	BB(*p)
	fmt.Println(p)
}

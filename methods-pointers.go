package main

import (
	"fmt"
	"math"
)

type V struct {
	x, y float64
}

func(v V) abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func (v V) scale(f float64) {
	v.x = v.x * f
	v.y = v.y * f
}

func main() {
	v := V{3,4}
	v.scale(10)
	fmt.Println(v.abs())
}

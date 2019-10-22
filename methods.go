package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	x, y float64
}

func abs(v Vertex) float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y + 100)
}

func (v Vertex) abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func main() {
	v := Vertex{3,4}
	fmt.Println(v.abs())

}
	

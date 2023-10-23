// Task 24

package main

import (
	"fmt"
	"math"
)

// Point struct with coordinates X and Y
type Point struct {
	X float64
	Y float64
}

// Point constructor
func newPoint(x, y float64) Point {
	p := Point{X: x, Y: y}
	return p
}

// Calculate and return distance
func GetDistance(a, b Point) float64 {
	return math.Sqrt(math.Pow((b.X-a.X), 2) + math.Pow((b.Y-a.Y), 2))
}

func main() {
	a := newPoint(1, -2.8)
	b := newPoint(-3, 6.6)
	fmt.Println(GetDistance(a, b))
}

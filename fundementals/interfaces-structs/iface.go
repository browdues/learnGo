package iface

import "math"

// Rectangle holds the data of a given rectangle
type Rectangle struct {
	Length float64
	Width  float64
}

// Circle blah
type Circle struct {
	radius float64
}

// Triangle blah
type Triangle struct {
	base   float64
	height float64
}

// Area method rectangle
func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

// Area method circle
func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

// Area method triangle
func (t Triangle) Area() float64 {
	return 0.5 * t.base * t.height
}

// Perimeter returns the distance around a giving rectangle
func Perimeter(rec Rectangle) float64 {
	return 2 * (rec.Length + rec.Width)
}

// Area returns the product of len x wid
func Area(rec Rectangle) float64 {
	return rec.Length * rec.Width
}

// Shape interface
type Shape interface {
	Area() float64
}

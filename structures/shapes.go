package structures

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Height float64
	Base   float64
}

type Shape interface {
	Area() float64
}

func Perimeter(rect Rectangle) float64 {
	return (rect.Height + rect.Width) * 2.0
}

func (rect Rectangle) Area() float64 {
	return rect.Height * rect.Width
}

func (circ Circle) Area() float64 {
	return math.Pi * math.Pow(circ.Radius, 2)
}

func (tria Triangle) Area() float64 {
	return tria.Height * tria.Base * 0.5
}

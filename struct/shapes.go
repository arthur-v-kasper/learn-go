package main

import "math"

// Shape is implemented by anything that can tell us its Area.
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle has the dimensions of a rectangle.
type Rectangle struct {
	Width  float64
	Height float64
}

// Area returns the area of the rectangle.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter returns the perimeter of a rectangle.
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle struct representa um circulo
type Circle struct {
	Radius float64
}

// Area retornara a area do circulo
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

//Perimeter retorno o perimetro do circulo
func (c Circle) Perimeter() float64 {
	return 2 * (math.Pi * c.Radius * c.Radius)
}

// Triangle struct representa um triangulo
type Triangle struct {
	Base   float64
	Height float64
}

// Area retornara a area da struc triangulo
func (t Triangle) Area() float64 {
	return t.Base * t.Height
}

//Perimeter retorno o perimetro do trinagulo
func (t Triangle) Perimeter() float64 {
	return 2 * (t.Base * t.Height)
}

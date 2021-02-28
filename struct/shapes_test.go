package main

import (
	"testing"
)

func TestPerimeterTriangle(t *testing.T) {
	triangle := Triangle{6.0, 6.0}
	wait := 72.0

	got := Shape.Perimeter(triangle)

	if got != wait {
		t.Errorf("Deu ruim galera!!")
	}
}

func TestAllPerimeters(t *testing.T) {
	//Array of struct that can be testing (TableDrivenTests)
	structPerimeters := []struct {
		name         string
		shape        Shape
		hasPerimeter float64
	}{
		{name: "Rectangule", shape: Rectangle{Width: 5, Height: 5}, hasPerimeter: 20.0},
		{name: "Circle", shape: Circle{Radius: 5.0}, hasPerimeter: 157.07963267948966},
		{name: "Triangule", shape: Triangle{Base: 5, Height: 5}, hasPerimeter: 50.0},
	}

	for _, tp := range structPerimeters {
		t.Run(tp.name, func(t *testing.T) {
			got := tp.shape.Perimeter()
			if got != tp.hasPerimeter {
				t.Errorf("Deu ruim, queria %g mas recebi %g", tp.hasPerimeter, got)
			}
		})
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 6, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})

	}

}

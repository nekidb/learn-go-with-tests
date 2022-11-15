package main

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("perimeter of rectangle with two float64 sides", func(t *testing.T) {
	
		rectangle := Rectangle{10.0, 10.0}
		got := rectangle.Perimeter()
		want := 40.0
		
		if got != want {
			t.Errorf("got %.2f, want %.2f", got, want)
		}
	})
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name 	string
		shape 	Shape
		hasArea float64
	} {
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 13}, hasArea: 530.92915845667505730018673177424},
		{name: "Triangle", shape: Triangle{Side_1: 3, Side_2: 4, Side_3: 5}, hasArea: 6.0},
	}
	
	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %.2f, want %.2f", tt.shape, got, tt.hasArea)
			}
		})
	}
}




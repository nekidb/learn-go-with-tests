package main

import "math"

type Shape interface {
	Area() float64
}


type Rectangle struct {
	Width 	float64
	Height 	float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}


type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}


type Triangle struct {
	Side_1 float64
	Side_2 float64
	Side_3 float64
}

func (t Triangle) Area() float64 {
	halfPerimeter := (t.Side_1 + t.Side_2 + t.Side_3) / 2
	
	return math.Sqrt(halfPerimeter * (halfPerimeter - t.Side_1) * (halfPerimeter - t.Side_2) * (halfPerimeter - t.Side_3))
}
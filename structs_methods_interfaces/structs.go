package structsmethodsinterfaces

import "math"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width

}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Triangle struct {
	Base   float64
	Height float64
	SideA  float64
	SideB  float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) * 0.5
}

func (t Triangle) Perimeter() float64 {
	return t.Base + t.SideA + t.SideB
}

type Trapezoid struct {
	Base1  float64
	Base2  float64
	SideA  float64
	SideB  float64
	Height float64
}

func (t Trapezoid) Area() float64 {
	return ((t.Base1 + t.Base2) / 2) * t.Height
}

func (t Trapezoid) Perimeter() float64 {
	return t.Base1 + t.Base2 + t.SideA + t.SideB
}

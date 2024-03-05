package solid

import "fmt"

// Интерфейс для фигур
type Shape interface {
	Area() float64
}

// Структура квадрата
type Square struct {
	SideLength float64
}

func (s *Square) Area() float64 {
	return s.SideLength * s.SideLength
}

// Структура прямоугольника
type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

func CalculateArea(s Shape) float64 {
	return s.Area()
}

func L() {
	fmt.Print("[Liskov Substitution Principle]\n")
	square := &Square{SideLength: 5}
	rectangle := &Rectangle{Width: 3, Height: 4}

	fmt.Println("Area of square:", CalculateArea(square))
	fmt.Println("Area of rectangle:", CalculateArea(rectangle))
}

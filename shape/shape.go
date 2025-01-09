package shape

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float32
}

type Square struct {
	sideLenght float32
}

func (square Square) Area() float32 {
	return square.sideLenght * square.sideLenght
}

type Circle struct {
	radius float32
}

func (circle Circle) Area() float32 {
	return circle.radius * circle.radius * math.Pi
}

func main() {
	square := Square{5}
	circle := Circle{8}
	printShapeArea(square)
	printShapeArea(circle)
}

func printShapeArea(shape Shape) {
	fmt.Println(shape.Area())
}

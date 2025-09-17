package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

// Создаем метод-конструктор для Point
func NewPoint(x float64, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func main() {
	//Объявляем и инициализируем точки
	point := NewPoint(1.1, 2.2)
	point1 := NewPoint(3.3, 4.4)

	fmt.Println(distance(point, point1))

}

// Создаем функцию и измеряем расстояние по формуле
func distance(point, point1 *Point) float64 {
	return math.Sqrt(math.Pow((point.x-point1.x), 2) + math.Pow((point.y-point1.y), 2))
}

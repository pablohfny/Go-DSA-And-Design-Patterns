package main

import "fmt"

type IDrawShape interface {
	drawShape(x [5]float32, y [5]float32)
}

type DrawShape struct {
}

func (shape DrawShape) drawShape(x [5]float32, y [5]float32) {
	fmt.Println("Drawing Shape")
}

type IContour interface {
	drawContour(x [5]float32, y [5]float32)
	resizeByFactor(factor int)
}

type DrawContour struct {
	x      [5]float32
	y      [5]float32
	shape  DrawShape
	factor int
}

func (contour DrawContour) resizeByFactor(factor int) {
	contour.factor = factor
}

func (contour DrawContour) drawContour(x [5]float32, y [5]float32) {
	fmt.Println("Drawing Contour")
	contour.shape.drawShape(contour.x, contour.y)
}

func main() {
	var x = [5]float32{1, 2, 3, 4, 5}
	var y = [5]float32{1, 2, 3, 4, 5}
	var contour IContour = DrawContour{x: x, y: y, factor: 2}

	contour.drawContour(x, y)
	contour.resizeByFactor(2)
}

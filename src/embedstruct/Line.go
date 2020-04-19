package main

type Point struct {
	x int
	y int
}

// type Point embedded into type Line
type Line struct {
	pointA Point
	pointB Point
}

func main()  {

	var line Line
	line.pointA.x = 0
	line.pointA.y = 0
	line.pointB.x = 9
	line.pointB.y = 9
}
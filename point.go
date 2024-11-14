package polygon

import "math"

type Point struct {
	X float64
	Y float64
}

func ManhattanDistance(a Point, b Point) float64 {
	return math.Abs(a.X-b.X) + math.Abs(a.Y-b.Y)
}

func EuclideanDistance(a Point, b Point) float64 {
	return math.Abs(math.Sqrt((a.X-b.X)*(a.X-b.X) + (a.Y-b.Y)*(a.Y-b.Y)))
}

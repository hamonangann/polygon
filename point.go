package polygon

import (
	"math"
)

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

func OrientationTriplet(a Point, b Point, c Point) int {
	orient := (b.Y-a.Y)*(c.X-b.X) - (b.X-a.X)*(c.Y-b.Y)

	if math.Abs(orient) < 1e-9 {
		return 0
	}

	if orient > 0 {
		return 1
	}

	return -1
}

func BetweenSegment(a Point, p Point, b Point) bool {
	return (OrientationTriplet(a, p, b) == 0) && ((p.X <= a.X) != (p.X <= b.X)) && ((p.Y <= a.Y) != (p.Y <= b.Y))
}

func LineSegmentIntersect(p1 Point, p2 Point, q1 Point, q2 Point) bool {
	o1 := OrientationTriplet(p1, p2, q1)
	o2 := OrientationTriplet(p1, p2, q2)
	o3 := OrientationTriplet(q1, q2, p1)
	o4 := OrientationTriplet(q1, q2, p2)

	if o1 != o2 && o3 != o4 {
		return true
	}

	b1 := BetweenSegment(p1, q1, p2)
	b2 := BetweenSegment(p1, q2, p2)
	b3 := BetweenSegment(q1, p1, q2)
	b4 := BetweenSegment(q1, p2, q2)

	return o1 == 0 && (b1 || b2 || b3 || b4)
}

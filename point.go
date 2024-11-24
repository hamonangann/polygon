package polygon

import (
	"math"
)

// Point is a type that stores a location of a point in a Cartesian coordinate system.
type Point struct {
	X float64
	Y float64
}

// ManhattanDistance returns the manhattan distance between two points.
// To measure the length of a line segment starts in a and ends in b, use EuclideanDistance instead.
func ManhattanDistance(a Point, b Point) float64 {
	return math.Abs(a.X-b.X) + math.Abs(a.Y-b.Y)
}

// EuclideanDistance returns the euclidean distance between two points.
func EuclideanDistance(a Point, b Point) float64 {
	return math.Abs(math.Sqrt((a.X-b.X)*(a.X-b.X) + (a.Y-b.Y)*(a.Y-b.Y)))
}

// OrientationTriplet returns an integer denoting the orientation made of triplet a->b->c.
//
// There are three possible outputs. If a->b->c is forming a clockwise triangle, it returns 1.
// If a->b->c is forming a counter-lockwise triangle, it returns -1.
// If a->b->c is collinear (forming a line), it returns 0.
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

// BetweenSegment returns true if Point p is located in a line segment from Point a to Point b.
// If a->p->b is not collinear, it always returns false.
func BetweenSegment(a Point, p Point, b Point) bool {
	return (OrientationTriplet(a, p, b) == 0) && ((p.X <= a.X) != (p.X <= b.X)) && ((p.Y <= a.Y) != (p.Y <= b.Y))
}

// LineSegmentIntersect returns true if a line segment from p1 to p2 intersects with a line segment from q1 to q2.
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

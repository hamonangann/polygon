package polygon

import (
	"errors"
	"math"
)

var (
	ErrInsufficientPoints = errors.New("points must be more than 2")          // ErrInsufficientPoints is for lack of Polygon points.
	ErrCollinearPoints    = errors.New("points must not collinear")           // ErrCollinearPoints is for Polygon with collinear points
	ErrSelfIntersecting   = errors.New("sides must not intersect each other") // ErrSelfIntersecting is for self-intersecting Polygon (not supported)
)

// Polygon is a type that stores a sequence of connected Point(s) forming valid simple polygon.
// It is immutable and all the points must be defined in the creation.
//
// ALWAYS construct Polygon(s) from NewPolygon() constructor or any utility functions in examples.go.
// Directly constructing Polygon{} may resulted in unexpected behavior.
type Polygon struct {
	points []Point
}

// Points returns a sequence of Point(s) forming the polygon.
func (p Polygon) Points() []Point {
	return p.points
}

// NewPolygon returns a pointer to Polygon (which is nil in case of error) and an error
// Error may happen if points provided to this constructor are forming an invalid simple polygon.
//
// There are three checks in this constructor.
// First, it checks for insufficient points. If less than 3 points are provided, they are not enough to form a polygon.
// Second, it checks for any collinearity by taking three neighboring points. It is considered invalid in this library.
// Third, if the polygon has more than 3 points (quadrilateral or more), it checks for self intersection.
// This library does not support self-intersecting polygons, since it is not a simple polygon anymore.
func NewPolygon(pts ...Point) (*Polygon, error) {
	if len(pts) < 3 {
		return nil, ErrInsufficientPoints
	}

	for i := range pts {
		for j := i + 1; j < len(pts)-1; j++ {
			if OrientationTriplet(pts[i], pts[j], pts[j+1]) == 0 {
				return nil, ErrCollinearPoints
			}
		}
	}

	if len(pts) > 3 {
		for i := range pts {
			for j := i + 2; j < len(pts); j++ {
				lineAStart := pts[i]
				lineAEnd := pts[i+1]
				lineBStart := pts[j]
				lineBEnd := pts[(j+1)%len(pts)]

				if lineAStart != lineBEnd && LineSegmentIntersect(lineAStart, lineAEnd, lineBStart, lineBEnd) {
					return nil, ErrSelfIntersecting
				}
			}
		}
	}

	return &Polygon{pts}, nil
}

// Sides returns the length of each sides of a Polygon
func (p Polygon) Sides() []float64 {
	sides := make([]float64, 0)

	for i := range p.points {
		if i == len(p.points)-1 {
			sides = append(sides, EuclideanDistance(p.points[i], p.points[0]))
		} else {
			sides = append(sides, EuclideanDistance(p.points[i], p.points[i+1]))
		}
	}

	return sides
}

// Perimeter returns the sum of length of each sides of a Polygon
func (p Polygon) Perimeter() float64 {
	ans := 0.0
	for i := range p.points {
		if i == len(p.points)-1 {
			ans += EuclideanDistance(p.points[i], p.points[0])
		} else {
			ans += EuclideanDistance(p.points[i], p.points[i+1])
		}
	}

	return ans
}

// Area returns the area inside a Polygon's sides.
func (p Polygon) Area() float64 {
	ans := 0.0
	for i := range p.points {
		if i == len(p.points)-1 {
			ans += p.points[i].X*p.points[0].Y - p.points[0].X*p.points[i].Y
		} else {
			ans += p.points[i].X*p.points[i+1].Y - p.points[i+1].X*p.points[i].Y
		}
	}

	return math.Abs(ans / 2)
}

// NumberOfSides returs the number of sides of a Polygon.
func (p Polygon) NumberOfSides() int {
	return len(p.points)
}

// NumberOfSides returs the number of vertices of a Polygon.
func (p Polygon) NumberOfVertices() int {
	return len(p.points)
}

// NumberOfSides returs the sum of interior angles of a Polygon in Radian.
func (p Polygon) SumOfInteriorAngles() float64 {
	return float64(len(p.points)-2) * math.Pi
}

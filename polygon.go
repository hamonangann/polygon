package polygon

import (
	"errors"
	"math"
)

var ErrInsufficientPoints = errors.New("points must be more than 2")
var ErrCollinearPoints = errors.New("points must not collinear")
var ErrSelfIntersecting = errors.New("sides must not intersect each other")

type Polygon struct {
	points []Point
}

func (p Polygon) Points() []Point {
	return p.points
}

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

func (p Polygon) NumberOfSides() int {
	return len(p.points)
}

func (p Polygon) NumberOfVertices() int {
	return len(p.points)
}

func (p Polygon) SumOfInteriorAngles() float64 {
	return float64(len(p.points)-2) * math.Pi
}

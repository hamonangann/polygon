package polygon

import (
	"errors"
	"math"
)

var ErrInsufficientPoints = errors.New("points must be more than 2")
var ErrCollinearPoints = errors.New("points must not be collinear")

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
			if Collinear(pts[i], pts[j], pts[j+1]) {
				return nil, ErrCollinearPoints
			}
		}

	}

	return &Polygon{pts}, nil
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

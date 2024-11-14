package polygon

import (
	"errors"
	"math"
)

var ErrInvalidPolygon = errors.New("points must be more than 2")

type Polygon struct {
	points []Point
}

func (p Polygon) Points() []Point {
	return p.points
}

func NewPolygon(pts ...Point) (*Polygon, error) {
	if len(pts) < 3 {
		return nil, ErrInvalidPolygon
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

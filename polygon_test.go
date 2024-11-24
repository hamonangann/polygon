package polygon_test

import (
	"math"
	"testing"

	"github.com/hamonangann/polygon"
	"github.com/stretchr/testify/assert"
)

func TestPolygon(t *testing.T) {
	t.Run("should handle polygon creation with less than 3 points", func(t *testing.T) {
		_, err := polygon.NewPolygon(
			polygon.Point{0, 0},
			polygon.Point{1, 2},
		)

		assert.EqualError(t, err, polygon.ErrInsufficientPoints.Error())
	})

	t.Run("should handle polygon creation with collinear points", func(t *testing.T) {
		_, err := polygon.NewPolygon(
			polygon.Point{1, 3},
			polygon.Point{2, 6},
			polygon.Point{3, 9},
		)

		assert.EqualError(t, err, polygon.ErrCollinearPoints.Error())
	})

	t.Run("should handle polygon creation with self-intersecting sides", func(t *testing.T) {
		_, err := polygon.NewPolygon(
			polygon.Point{0, 3},
			polygon.Point{3, 3},
			polygon.Point{0, 0},
			polygon.Point{3, 0},
		)

		assert.EqualError(t, err, polygon.ErrSelfIntersecting.Error())
	})

	t.Run("should return created points from valid polygon", func(t *testing.T) {
		pts := []polygon.Point{
			{0, 0},
			{4, 3},
			{4, 0},
		}
		p, _ := polygon.NewPolygon(pts...)

		assert.Equal(t, pts, p.Points())
	})

	t.Run("should return correct sides length", func(t *testing.T) {
		p, _ := polygon.NewPolygon(
			polygon.Point{0, 0},
			polygon.Point{0, 3},
			polygon.Point{4, 3},
			polygon.Point{4, 0},
		)

		assert.Equal(t, []float64{3, 4, 3, 4}, p.Sides())
	})

	t.Run("should return correct perimeter", func(t *testing.T) {
		p, _ := polygon.NewPolygon(
			polygon.Point{0, 0},
			polygon.Point{0, 3},
			polygon.Point{4, 3},
			polygon.Point{4, 0},
		)

		assert.Equal(t, 14.0, p.Perimeter())
	})

	t.Run("should return correct area", func(t *testing.T) {
		p, _ := polygon.NewPolygon(
			polygon.Point{0, 0},
			polygon.Point{0, 3},
			polygon.Point{4, 0},
		)

		assert.Equal(t, 6.0, p.Area())
	})

	t.Run("should return correct number of sides", func(t *testing.T) {
		p, _ := polygon.NewPolygon(
			polygon.Point{0, 0},
			polygon.Point{0, 3},
			polygon.Point{3, 4},
			polygon.Point{5, 3},
			polygon.Point{5, 0},
		)

		assert.Equal(t, 5, p.NumberOfSides())
	})

	t.Run("should return correct number of vertices", func(t *testing.T) {
		p, _ := polygon.NewPolygon(
			polygon.Point{0, 0},
			polygon.Point{0, 3},
			polygon.Point{3, 4},
			polygon.Point{5, 3},
			polygon.Point{5, 0},
		)

		assert.Equal(t, 5, p.NumberOfVertices())
	})

	t.Run("should return correct sum of interior angles", func(t *testing.T) {
		p, _ := polygon.NewPolygon(
			polygon.Point{0, 0},
			polygon.Point{0, 3},
			polygon.Point{3, 4},
			polygon.Point{5, 3},
			polygon.Point{5, 0},
		)

		assert.Equal(t, math.Pi*3.0, p.SumOfInteriorAngles())
	})
}

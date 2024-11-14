package polygon_test

import (
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

		assert.EqualError(t, err, polygon.ErrInvalidPolygon.Error())
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
}

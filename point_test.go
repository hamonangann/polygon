package polygon_test

import (
	"testing"

	"github.com/hamonangann/polygon"
	"github.com/stretchr/testify/assert"
)

func TestPoint(t *testing.T) {
	t.Run("should return correct manhattan distance", func(t *testing.T) {
		a := polygon.Point{0, 0}
		b := polygon.Point{3, 4}

		assert.Equal(t, 7.0, polygon.ManhattanDistance(a, b))
	})

	t.Run("should return correct euclidean distance", func(t *testing.T) {
		a := polygon.Point{0, 0}
		b := polygon.Point{3, 4}

		assert.Equal(t, 5.0, polygon.EuclideanDistance(a, b))
	})

	t.Run("should return 1 if a triplet of points is clockwise", func(t *testing.T) {
		a := polygon.Point{0, 0}
		b := polygon.Point{3, 3}
		c := polygon.Point{3, 0}

		assert.Equal(t, 1, polygon.OrientationTriplet(a, b, c))
	})

	t.Run("should return -1 if a triplet of points is counterclockwise", func(t *testing.T) {
		a := polygon.Point{0, 0}
		b := polygon.Point{3, 3}
		c := polygon.Point{0, 3}

		assert.Equal(t, -1, polygon.OrientationTriplet(a, b, c))
	})

	t.Run("should return true if a triplet of points is collinear", func(t *testing.T) {
		a := polygon.Point{1, 3}
		b := polygon.Point{0, 0}
		c := polygon.Point{-1, -3}

		assert.Equal(t, 0, polygon.OrientationTriplet(a, b, c))
	})

	t.Run("should return true if a point is between two points forming line segment", func(t *testing.T) {
		a := polygon.Point{1, 3}
		b := polygon.Point{0, 0}
		c := polygon.Point{-1, -3}

		assert.Equal(t, true, polygon.BetweenSegment(a, b, c))
	})

	t.Run("should return true if two line segments intersect each other", func(t *testing.T) {
		a1 := polygon.Point{0, 0}
		a2 := polygon.Point{3, 3}
		b1 := polygon.Point{3, 0}
		b2 := polygon.Point{0, 3}

		assert.Equal(t, true, polygon.LineSegmentIntersect(a1, a2, b1, b2))
	})

	t.Run("should return true if there are collinearity between two line segments and the segements intersect, case 1", func(t *testing.T) {
		a1 := polygon.Point{0, 0}
		a2 := polygon.Point{2, 2}
		b1 := polygon.Point{1, 1}
		b2 := polygon.Point{3, 3}

		assert.Equal(t, true, polygon.LineSegmentIntersect(a1, a2, b1, b2))
	})

	t.Run("should return true if there are collinearity between two line segments and the segements intersect, case 2", func(t *testing.T) {
		a1 := polygon.Point{3, 3}
		a2 := polygon.Point{4, 4}
		b1 := polygon.Point{1, 1}
		b2 := polygon.Point{3, 3}

		assert.Equal(t, true, polygon.LineSegmentIntersect(a1, a2, b1, b2))
	})

	t.Run("should return true if there are collinearity between two line segments and the segements intersect, case 3", func(t *testing.T) {
		a1 := polygon.Point{4, 4}
		a2 := polygon.Point{2, 2}
		b1 := polygon.Point{2, 2}
		b2 := polygon.Point{1, 1}

		assert.Equal(t, true, polygon.LineSegmentIntersect(a1, a2, b1, b2))
	})

	t.Run("should return true if there are collinearity between two line segments and the segements intersect, case 4", func(t *testing.T) {
		a1 := polygon.Point{1, 1}
		a2 := polygon.Point{4, 4}
		b1 := polygon.Point{5, 5}
		b2 := polygon.Point{2, 2}

		assert.Equal(t, true, polygon.LineSegmentIntersect(a1, a2, b1, b2))
	})
}

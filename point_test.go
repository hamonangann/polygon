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
}

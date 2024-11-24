package polygon_test

import (
	"testing"

	"github.com/hamonangann/polygon"
	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	t.Run("should return correct triangle", func(t *testing.T) {
		pl := polygon.NewTriangle(3, 4)

		assert.Equal(t, 6.0, pl.Area())
		assert.Equal(t, 12.0, pl.Perimeter())
	})

	t.Run("should return correct square", func(t *testing.T) {
		pl := polygon.NewSquare(3)

		assert.Equal(t, 9.0, pl.Area())
		assert.Equal(t, 12.0, pl.Perimeter())
	})

	t.Run("should return correct rectangle", func(t *testing.T) {
		pl := polygon.NewRectangle(3, 4)

		assert.Equal(t, 12.0, pl.Area())
		assert.Equal(t, 14.0, pl.Perimeter())
	})

	t.Run("should return correct rhombus", func(t *testing.T) {
		pl := polygon.NewRhombus(6, 8)

		assert.Equal(t, 24.0, pl.Area())
		assert.Equal(t, 20.0, pl.Perimeter())
	})

	t.Run("should return correct trapezoid", func(t *testing.T) {
		pl := polygon.NewTrapezoid(7, 4, 4)

		assert.Equal(t, 22.0, pl.Area())
		assert.Equal(t, 20.0, pl.Perimeter())
	})
}

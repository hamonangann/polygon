package polygon

func NewTriangle(base float64, height float64) Polygon {
	pl, _ := NewPolygon(
		Point{0, 0},
		Point{base, 0},
		Point{0, height},
	)

	return *pl
}

func NewSquare(side float64) Polygon {
	pl, _ := NewPolygon(
		Point{0, 0},
		Point{0, side},
		Point{side, side},
		Point{side, 0},
	)

	return *pl
}

func NewRectangle(length float64, height float64) Polygon {
	pl, _ := NewPolygon(
		Point{0, 0},
		Point{length, 0},
		Point{length, height},
		Point{0, height},
	)

	return *pl
}

func NewRhombus(diagonalA float64, diagonalB float64) Polygon {
	pl, _ := NewPolygon(
		Point{diagonalA / 2, 0},
		Point{0, diagonalB / 2},
		Point{diagonalA / 2, diagonalB},
		Point{diagonalA, diagonalB / 2},
	)

	return *pl
}

func NewTrapezoid(sideA float64, sideB float64, height float64) Polygon {
	pl, _ := NewPolygon(
		Point{0, 0},
		Point{0, sideA},
		Point{sideB, height},
		Point{height, 0},
	)

	return *pl
}

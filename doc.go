// Package polygon is a library for constructing [simple polygons].
// It provides methods to measure polygon properties, such as perimeter, area, and sum of internal angles.
// A simple polygon is a sequence of points that are connecting to each other without self-intersection.
//
// Euclidean geometry is assumed throughout this library.
//
// Example: (constructing right triangle with 3 points)
//
//	p, _ := polygon.NewPolygon(
//	    polygon.Point{0, 0},
//	    polygon.Point{0, 3},
//	    polygon.Point{4, 0},
//	)
//	fmt.Println(p.Area()) // 6.0
//
// [simple polygon]: https://en.wikipedia.org/wiki/Simple_polygon
package polygon

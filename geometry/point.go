package geometry

import (
	"fmt"
)

// Point identifies a point on the x/y plane
type Point struct {
	x Coordinate
	y Coordinate
}

// String returns a string representing the point
func (p Point) String() string {

	return fmt.Sprintf("(%s,%s)", p.x, p.y)
}

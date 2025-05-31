package geometry

import "fmt"

// Polygon  attribute to identify an x or y coordinate for a point. Should be at least 3 corners
type Polygon struct {
	corners []Point
}

// String returns a string representing the polygon from starting point to end point and back to starting point
func (poly Polygon) String() string {
	outputString := ""
	for _, c := range poly.corners {
		if outputString != "" {
			outputString += "-" // add - if we already started
		}
		outputString += c.String()
	}
	if len(poly.corners) > 0 {
		outputString += "-" + poly.corners[0].String() // add first point at the end again
	}
	return outputString
}

// NewPolygon is constructor for a polygon, should be at least 3 corners
func NewPolygon(cors []Point) (*Polygon, error) {
	if len(cors) < 3 {
		return &Polygon{}, fmt.Errorf("polygon needs at least 3 corners")
	}
	return &Polygon{corners: cors}, nil
}

// Edges returns the line segments that represent the polygon in order
func (poly *Polygon) Edges() []Line {
	output := []Line{}
	for i := range poly.corners { // going over indices only
		if i != 0 {
			output = append(output, Line{startingPoint: poly.corners[i-1], endPoint: poly.corners[i]})
		}
	}
	// add last segment from last point to first point
	output = append(output, Line{startingPoint: poly.corners[len(poly.corners)-1], endPoint: poly.corners[0]})
	return output
}

// IsPointInside returns  whether a point is inside the polygon or not. Uses the ray tracing method by drawing a line from the point to INF on x axis and
// checking whether it intersects even number of edges (outside), or odd number of edges (inside)
func (poly *Polygon) IsPointInside(p Point) bool {
	isInside := false
	insideRay := Line{startingPoint: p, endPoint: Point{INF, p.y}} // ray from point to be tested to INF on x-axis
	for _, e := range poly.Edges() {                               // going over all edges
		if insideRay.DoesIntersect(e) {
			isInside = !isInside
		}
	}
	return isInside
}

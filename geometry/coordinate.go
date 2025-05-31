package geometry

import (
	"fmt"
	"math"
)

// Coordinate attribute to identify an x or y coordinate for a point
type Coordinate float32

const (
	// INF constant to identify Positive INF, set to max positive float32 number
	INF Coordinate = math.MaxFloat32

	// NEG_INF constant to identify Positive INF, set to max positive float32 number
	NEG_INF Coordinate = -1 * math.MaxFloat32

	// NEG_INF constant to identify Positive INF, set to max positive float32 number
	MINUS_INF Coordinate = NEG_INF
)

// String returns a string representing the coordinate as a float
func (c Coordinate) String() string {
	if c == INF {
		return "INF"
	} else if c == NEG_INF {
		return "-INF"
	}
	return fmt.Sprintf("%g", c)
}

// InRange is a helper function to determine if a coordinate is in range of 2 coordinates. Starting coordinate counts, ending coordinate doesn't count
func (c Coordinate) InRange(sc Coordinate, ec Coordinate) bool {
	if sc < ec { // start is before end
		if ec == INF {
			if c >= sc && c <= ec {
				return true
			} else {
				return false
			}
		} else {
			if c >= sc && c < ec {
				return true
			} else {
				return false
			}
		}

	} else if sc > ec {
		if ec == NEG_INF {
			if c >= ec && c <= sc {
				return true
			} else {
				return false
			}
		} else {
			if c > ec && c <= sc {
				return true
			} else {
				return false
			}
		}
	} else { // sc==ec
		if c == sc {
			return true
		} else {
			return false
		}
	}

}

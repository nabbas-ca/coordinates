package geometry

import (
	"fmt"
)

// Line is a struct to identify an line segment in x/y plane
type Line struct {
	startingPoint Point
	endPoint      Point
}

// String returns a string representing the line
func (l Line) String() string {

	return fmt.Sprintf("%s-%s", l.startingPoint, l.endPoint)
}

// Inverse return an inverse of the line
/*func (l Line) Inverse() Line {
	return Line{startingPoint: l.endPoint, endPoint: l.startingPoint}
}*/

// Inverse return an inverse of the line
func (l Line) GoingLeftHorizontally() bool {
	return l.startingPoint.x > l.endPoint.x // if starting point is higher in x value, return true, otherwise return false

}

// OrientationProduct is the orientation product of a point with respect to a line to determine the orientation of a point with respect to the line.
//
//	If negative then clockwise, positive then counter clockwise. The output in float64 to handle infinity
func (l Line) OrientationProduct(p Point) float64 {

	//the vectors are l.endpoint to p and l.starting point to l.endpoint
	// x1,y1 is l.startingpoint
	// x2,y2 is l.endpoint
	// x3,y3 is p
	// orientation product according to this AI output: (x3-x2)(y2-y1)-(x2-x1)(y3-y2)
	// which is the product of | (x3-x2)  (x2-x1) |
	//                         | (y3-y2)  (y2-y1) |
	// What's given in the problem      Two line segments are defined by their endpoints. Line segment 1: \((x_{1},y_{1})\) to \((x_{2},y_{2})\). Line segment 2: \((x_{3},y_{3})\) to \((x_{4},y_{4})\).         Helpful information      The orientation of three points can be used to determine if two line segments intersect. The orientation can be clockwise, counterclockwise, or collinear. The formula for orientation is: \((y_{2}-y_{1})(x_{3}-x_{2})-(x_{2}-x_{1})(y_{3}-y_{2})\). If the orientation is positive, the points are counterclockwise. If the orientation is negative, the points are clockwise. If the orientation is zero, the points are collinear.  .f5cPye .WaaZC:first-of-type .rPeykc.uP58nb:first-child{font-size:var(--m3t3);line-height:var(--m3t4);font-weight:400 !important;letter-spacing:normal;margin:0 0 10px 0}.rPeykc.uP58nb{font-size:var(--m3t5);font-weight:500;letter-spacing:0;line-height:var(--m3t6);margin:20px 0 10px 0}.rPeykc.uP58nb.MNX06c{font-size:var(--m3t1);font-weight:normal;letter-spacing:normal;line-height:var(--m3t2);margin:10px 0 10px 0}.f5cPye ol{font-size:var(--m3t7);line-height:var(--m3t8);margin:10px 0 20px 0;padding-left:24px}.f5cPye .WaaZC:first-of-type ol:first-child{margin-top:0}.f5cPye ol.qh1nvc{font-size:var(--m3t7);line-height:var(--m3t8)}.PpKptb{color:var(--m3c11) !important;font-family:Google Sans,Arial,sans-serif;font-size:var(--m3t11);font-weight:500;line-height:var(--m3t12)}.BFxDoe{color:var(--m3c10) !important;font-family:Google Sans,Arial,sans-serif;font-size:var(--m3t9);letter-spacing:0.1px;line-height:var(--m3t10)}.UnzV3b{color:var(--m3c11);font-size:var(--m3t7);line-height:var(--m3t8)}.f5cPye ul .UrtGC,.f5cPye ol .UrtGC{margin-left:-24px}.UrtGC .dnXCYb[aria-expanded="true"] .WltAjf,.UrtGC .dnXCYb.yMbVTb .WltAjf{-webkit-line-clamp:unset}.UrtGC .dnXCYb{overflow:hidden}.aj35ze{fill:#747878;display:inline-block;height:24px;width:24px}.h373nd{position:relative}.dnXCYb{align-items:center;box-sizing:border-box;display:flex;position:relative;width:100%;cursor:pointer}html:not(.zAoYTe) .dnXCYb{outline:0}.JlqpRe{flex:1;margin:12px 0;overflow:hidden}.h373nd:not(.LJm5W) .JCzEY{font-weight:500}.TQc1id .ABs8Y{font-weight:500}.ABs8Y,.JCzEY{color:#1f1f1f}.APjcId,.WltAjf{color:var(--IXoxUe)}.WltAjf::before{content:'';display:block;height:4px}.bCOlv{width:100%}.bCOlv:not(.yMbVTb){position:absolute;display:none;opacity:0}.bCOlv:not(.yMbVTb) .GKFAcc{opacity:0}.IZE3Td{position:relative}.ru2Kjc{display:none}.L3Ezfd{position:absolute;height:100%;width:100%;left:0;top:0}.J2MhIb.LJm5W .JCzEY{font-weight:700}.ABs8Y,.JCzEY,.bJi8Dd,.APjcId,.WltAjf{display:-webkit-box;-webkit-box-orient:vertical;overflow:hidden}.JCzEY{-webkit-line-clamp:2}.gVe2qd{-webkit-line-clamp:unset !important;word-break:unset !important}.yMbVTb.dnXCYb .aj35ze{transform:scale3d(1,-1,1)}.iXPZfd.dnXCYb .ABs8Y,.iXPZfd.dnXCYb .JCzEY{-webkit-line-clamp:unset !important;word-break:unset !important}.yMbVTb.dhks6d .APjcId,.yMbVTb.dhks6d .WltAjf{opacity:0;height:0}.APjcId,.WltAjf{-webkit-line-clamp:1}.CC4Ctb .JCzEY{-webkit-line-clamp:1;word-break:break-all}.LJm5W .CC4Ctb.dnXCYb{min-height:calc(40px + 2*12px)}.ilulF .ABs8Y,.ilulF .JCzEY,.ilulF .APjcId,.ilulF .WltAjf{-webkit-line-clamp:unset!important;word-break:unset!important}.iRPzcb{border-bottom:1px solid var(--gS5jXb)}.iwY1Mb{height:0;width:0;opacity:0;display:block}.fxvkXe,.p8Jhnd{width:36px;height:36px;background:#f1f3f4;border-radius:50%;display:flex;justify-content:center;align-items:center;flex-shrink:0;margin:0 0 0 12px}.dnXCYb:not(.FjLqqd):not(.CC4Ctb) .p8Jhnd{margin:12px 0 12px 12px}       How to solve             Calculate the orientation of the endpoints of each line segment with respect to the other line segment.      Step 1 . Calculate the orientation of point \((x_{3},y_{3})\) with respect to the line segment 1. \(o_{1}=(y_{2}-y_{1})(x_{3}-x_{2})-(x_{2}-x_{1})(y_{3}-y_{2})\) Step 2 . Calculate the orientation of point \((x_{4},y_{4})\) with respect to the line segment 1. \(o_{2}=(y_{2}-y_{1})(x_{4}-x_{2})-(x_{2}-x_{1})(y_{4}-y_{2})\) Step 3 . Calculate the orientation of point \((x_{1},y_{1})\) with respect to the line segment 2. \(o_{3}=(y_{4}-y_{3})(x_{1}-x_{4})-(x_{4}-x_{3})(y_{1}-y_{4})\) Step 4 . Calculate the orientation of point \((x_{2},y_{2})\) with respect to the line segment 2. \(o_{4}=(y_{4}-y_{3})(x_{2}-x_{4})-(x_{4}-x_{3})(y_{2}-y_{4})\) Step 5 . Check if the line segments intersect.

	//var output float64
	output := ((float64)(p.x-l.endPoint.x))*((float64)(l.endPoint.y-l.startingPoint.y)) - ((float64)(l.endPoint.x-l.startingPoint.x))*((float64)(p.y-l.endPoint.y))
	return output

}

// Sign is a helper function to determine orientation
func Sign(x float64) string {
	if x > 0 {
		return "+"
	}
	if x < 0 {
		return "-"
	}
	return "0"
}

// DoesIntersect
//
//	  returns true if l2 intersects l1, starting and endpoints don't count in the calculation.
//		   for example, (0,0)-(2-0) does not intersect with (1,0)-(1,1)
func (l1 Line) DoesIntersect(l2 Line) bool {

	// Full algorithm:
	// determine orientation of l2 endpoints to l1, they should be opposite and not collinear
	// determine orientation of l1 endpoints to l2, they should be opposite and not collinear

	// Part1: determine orientation of l2 endpoints to l1, they should be opposite and not collinear
	o1 := Sign(l1.OrientationProduct(l2.startingPoint))
	o2 := Sign(l1.OrientationProduct(l2.endPoint))

	if o1 == "0" {
		if l2.startingPoint.x.InRange(l1.startingPoint.x, l1.endPoint.x) && l2.startingPoint.y.InRange(l1.startingPoint.y, l1.endPoint.y) {
			return true
		} else {
			return false
		}
	}
	if o2 == "0" || o1 == o2 {
		return false
	}

	// Part2: deterdetermine orientation of l1 endpoints to l2, they should be opposite and not collinear
	o3 := Sign(l2.OrientationProduct(l1.startingPoint))
	o4 := Sign(l2.OrientationProduct(l1.endPoint))

	if o3 == "0" {
		if l1.startingPoint.x.InRange(l2.startingPoint.x, l2.endPoint.x) && l1.startingPoint.y.InRange(l2.startingPoint.y, l2.endPoint.y) {
			return true
		} else {
			return false
		}
	}
	if o4 == "0" || o3 == o4 {
		return false
	}

	return true // otherwise they intersect

	/*if l1.GoingLeftHorizontally() {
		return l1.Inverse().DoesIntersect(l2)
	}
	if l2.GoingLeftHorizontally() {
		return l1.DoesIntersect(l2.Inverse())
	}
	*/

	//return fmt.Sprintf("%s-%s", l.startingPoint, l.endPoint)
}

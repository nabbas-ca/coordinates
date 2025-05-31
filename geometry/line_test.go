package geometry

import (
	"testing"
)

func TestLine_String(t *testing.T) {

	tests := []struct {
		name string
		line Line
		want string
	}{
		{
			name: "(0,0) to (1,1)",
			line: Line{startingPoint: Point{x: 0.0, y: 0.0}, endPoint: Point{x: 1.0, y: 1.0}},
			want: "(0,0)-(1,1)",
		},
		{
			name: "(1,1) to (0,0)",
			line: Line{startingPoint: Point{x: 1.0, y: 1.0}, endPoint: Point{x: 0, y: 0}},
			want: "(1,1)-(0,0)",
		},
		{
			name: "(1,1) to (1,INF)",
			line: Line{startingPoint: Point{x: 1.0, y: 1.0}, endPoint: Point{x: 1, y: INF}},
			want: "(1,1)-(1,INF)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.line.String(); got != tt.want {
				t.Errorf("Line.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLine_GoingLeftHorizontally(t *testing.T) {

	tests := []struct {
		name string
		line Line
		want bool
	}{
		{
			name: "(1,1) to (0,0) is left",
			line: Line{startingPoint: Point{x: 1.0, y: 1.0}, endPoint: Point{x: 0, y: 0}},
			want: true,
		},
		{
			name: "(0,0) to (1,1) is not left",
			line: Line{startingPoint: Point{x: 0.0, y: 0.0}, endPoint: Point{x: 1, y: 1}},
			want: false,
		},
		{
			name: "(0,0) to (-INF,0) is left",
			line: Line{startingPoint: Point{x: 0.0, y: 0.0}, endPoint: Point{x: NEG_INF, y: 0}},
			want: true,
		},
		{
			name: "(0,0) to (INF,0) is not left",
			line: Line{startingPoint: Point{x: 0.0, y: 0.0}, endPoint: Point{x: INF, y: 0}},
			want: false,
		},
		{
			name: "(0,0) to (0,INF) is vertical, not left",
			line: Line{startingPoint: Point{x: 0.0, y: 0.0}, endPoint: Point{x: 0, y: INF}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.line.GoingLeftHorizontally(); got != tt.want {
				t.Errorf("Line.GoingLeftHorizontally() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLine_OrientationProduct(t *testing.T) {

	tests := []struct {
		name string
		l    Line
		p    Point
		want string
	}{
		{
			name: "(0,0) - (1,1) orientation to (1,0) should be positive(clockwise)",
			l:    Line{startingPoint: Point{x: 0.0, y: 0.0}, endPoint: Point{x: 1.0, y: 1.0}},
			p:    Point{x: 1.0, y: 0.0},
			want: "+", // - means answer should be negative
		},
		{
			name: "(0,0) - (1,1) orientation to (0,1) should be negative (counter clockwise)",
			l:    Line{startingPoint: Point{x: 0.0, y: 0.0}, endPoint: Point{x: 1.0, y: 1.0}},
			p:    Point{x: 0, y: 1},
			want: "-", // - means answer should be negative
		},
		{
			name: "(0,0) - (1,1) orientation to (2,2) should be zero (collinear)",
			l:    Line{startingPoint: Point{x: 0.0, y: 0.0}, endPoint: Point{x: 1.0, y: 1.0}},
			p:    Point{x: 2, y: 2},
			want: "0", // - means answer should be negative
		},
		{
			name: "test inf : (0,0) - (1,1) orientation to (inf,2) should be positive(clockwise)",
			l:    Line{startingPoint: Point{x: 0.0, y: 0.0}, endPoint: Point{x: 1.0, y: 1.0}},
			p:    Point{x: INF, y: 2},
			want: "+", // - means answer should be negative
		},
		{
			name: "test -inf : (0,0) - (1,1) orientation to (-inf,2) should be negative(counter clockwise)",
			l:    Line{startingPoint: Point{x: 0.0, y: 0.0}, endPoint: Point{x: 1.0, y: 1.0}},
			p:    Point{x: NEG_INF, y: 2},
			want: "-", // - means answer should be negative
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			if got := tt.l.OrientationProduct(tt.p); Sign(got) != tt.want {
				t.Errorf("Line.OrientationProduct() = %v, Sign(got)=%s, want %v", got, Sign(got), tt.want)
			}

		})
	}
}

func TestLine_DoesIntersect(t *testing.T) {

	tests := []struct {
		name string
		l1   Line
		l2   Line
		want bool
	}{
		{
			name: "(0,0) - (1,1) should intersect with (0,1)-(1,0) [like an X]",
			l1:   Line{startingPoint: Point{x: 0.0, y: 0.0}, endPoint: Point{x: 1.0, y: 1.0}},
			l2:   Line{startingPoint: Point{x: 0.0, y: 1.0}, endPoint: Point{x: 1.0, y: 0}},
			want: true,
		},
		{
			name: "(0,0) - (1,0) should not intersect with (0,1)-(1,1) [they are parallel horizontally]",
			l1:   Line{startingPoint: Point{x: 0.0, y: 0.0}, endPoint: Point{x: 1.0, y: 0}},
			l2:   Line{startingPoint: Point{x: 0.0, y: 1.0}, endPoint: Point{x: 1.0, y: 1}},
			want: false,
		},
		{
			name: "(0,0) - (2,0) should  intersect with (1,0)-(1,1) [end points don't count, starting point counts]",
			l1:   Line{startingPoint: Point{x: 0.0, y: 0.0}, endPoint: Point{x: 2.0, y: 0}},
			l2:   Line{startingPoint: Point{x: 1, y: 0}, endPoint: Point{x: 1, y: 1}},
			want: true,
		},
		{
			name: "(0,0) - (2,0) should not intersect with (1,1)-(1,0) [end points don't count, starting point counts]",
			l1:   Line{startingPoint: Point{x: 0.0, y: 0.0}, endPoint: Point{x: 2.0, y: 0}},
			l2:   Line{startingPoint: Point{x: 1, y: 1}, endPoint: Point{x: 1, y: 0}},
			want: false,
		},
		{
			name: "(0,0) - (2,0) should  intersect with (1,-1)-(1,1) [like a + sign]",
			l1:   Line{startingPoint: Point{x: 0.0, y: 0.0}, endPoint: Point{x: 2.0, y: 0}},
			l2:   Line{startingPoint: Point{x: 1, y: -1}, endPoint: Point{x: 1, y: 1}},
			want: true,
		},
		{
			name: "(0,0)- (INF,0) should intersect with (1,-1)-(1,1) [like a looooong + sign]",
			l1:   Line{startingPoint: Point{x: 0.0, y: 0.0}, endPoint: Point{x: INF, y: 0}},
			l2:   Line{startingPoint: Point{x: 1, y: -1}, endPoint: Point{x: 1, y: 1}},
			want: true,
		},
		{
			name: "(-INF,0) - (0,0) should  not intersect with (1,-1)-(1,1) ",
			l1:   Line{startingPoint: Point{x: NEG_INF, y: 0.0}, endPoint: Point{x: 0, y: 0}},
			l2:   Line{startingPoint: Point{x: 1, y: -1}, endPoint: Point{x: 1, y: 1}},
			want: false,
		},
		{
			name: "(INF,0) - (0,0) should  not intersect with (0,-1)-(0,1) [end points don't count]",
			l1:   Line{startingPoint: Point{x: INF, y: 0.0}, endPoint: Point{x: 0, y: 0}},
			l2:   Line{startingPoint: Point{x: 0, y: -1}, endPoint: Point{x: 0, y: 1}},
			want: false,
		},
		{
			name: "(-INF,0) - (0,0) should  not intersect with (0,-1)-(0,1) [end points don't count]",
			l1:   Line{startingPoint: Point{x: NEG_INF, y: 0.0}, endPoint: Point{x: 0, y: 0}},
			l2:   Line{startingPoint: Point{x: 0, y: -1}, endPoint: Point{x: 0, y: 1}},
			want: false,
		},
		{
			name: "(0,0)-(INF,0) should not intersect with (1,1)-(1,3)",
			l1:   Line{startingPoint: Point{x: 0.0, y: 0.0}, endPoint: Point{x: INF, y: 0}},
			l2:   Line{startingPoint: Point{x: 1, y: 1}, endPoint: Point{x: 1, y: 3}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.l1.DoesIntersect(tt.l2); got != tt.want {
				t.Errorf("l1.Line.DoesIntersect(l2) = %v, want %v", got, tt.want)
			}
			// test reverse
			if got := tt.l2.DoesIntersect(tt.l1); got != tt.want {
				t.Errorf("l2.DoesIntersect(l1) = %v, want %v", got, tt.want)
			}
			// test inverses
			/*
				if got := tt.l1.Inverse().DoesIntersect(tt.l2); got != tt.want {
					t.Errorf("Line.DoesIntersect() = %v, want %v", got, tt.want)
				}
				if got := tt.l1.DoesIntersect(tt.l2.Inverse()); got != tt.want {
					t.Errorf("Line.DoesIntersect() = %v, want %v", got, tt.want)
				}
				if got := tt.l1.Inverse().DoesIntersect(tt.l2.Inverse()); got != tt.want {
					t.Errorf("Line.DoesIntersect() = %v, want %v", got, tt.want)
				}*/
		})
	}
}

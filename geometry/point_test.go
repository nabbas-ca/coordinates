package geometry

import "testing"

func TestPoint_String(t *testing.T) {

	tests := []struct {
		name string
		p    Point
		want string
	}{
		{
			name: "test 0,0",
			p:    Point{x: 0.0, y: 0.0},
			want: "(0,0)",
		},
		{
			name: "test 0,+INF",
			p:    Point{x: 0.0, y: INF},
			want: "(0,INF)",
		},
		{
			name: "test 0,NEG_INF",
			p:    Point{x: 0.0, y: NEG_INF},
			want: "(0,-INF)",
		},
		{
			name: "test +INF,0",
			p:    Point{x: INF, y: 0.0},
			want: "(INF,0)",
		},
		{
			name: "test NEG_INF,0",
			p:    Point{x: NEG_INF, y: 0.0},
			want: "(-INF,0)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.p.String(); got != tt.want {
				t.Errorf("Point.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

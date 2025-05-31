package geometry

import (
	"reflect"
	"testing"
)

func TestPolygon_String(t *testing.T) {
	tests := []struct {
		name string
		poly Polygon
		want string
	}{
		{
			name: "triangle: (0,0)-(1,1)-(1,0)-(0,0)",
			poly: Polygon{corners: []Point{
				{x: 0, y: 0},
				{x: 1, y: 1},
				{x: 1, y: 0},
			},
			},
			want: "(0,0)-(1,1)-(1,0)-(0,0)",
		},
		{
			name: "square: (0,0)-(0,1)-(1,1)-(1,0)-(0,0)",
			poly: Polygon{corners: []Point{
				{x: 0, y: 0},
				{x: 0, y: 1},
				{x: 1, y: 1},
				{x: 1, y: 0},
			},
			},
			want: "(0,0)-(0,1)-(1,1)-(1,0)-(0,0)",
		},
		{
			name: "4 star polygon: (0,3)-(1,1)-(3,0)-(1,-1)-(0,-3)-(-1,-1)-(-3,0)-(-1,1)",
			poly: Polygon{corners: []Point{
				{x: 0, y: 3},
				{x: 1, y: 1},
				{x: 3, y: 0},
				{x: 1, y: -1},
				{x: 0, y: -3},
				{x: -1, y: -1},
				{x: -3, y: 0},
				{x: -1, y: 1},
			},
			},
			want: "(0,3)-(1,1)-(3,0)-(1,-1)-(0,-3)-(-1,-1)-(-3,0)-(-1,1)-(0,3)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.poly.String(); got != tt.want {
				t.Errorf("Polygon.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewPolygon(t *testing.T) {

	tests := []struct {
		name    string
		cors    []Point
		want    *Polygon
		wantErr bool
	}{
		{
			name: "4 star polygon: (0,3)-(1,1)-(3,0)-(1,-1)-(0,-3)-(-1,-1)-(-3,0)-(-1,1)",
			cors: []Point{
				{x: 0, y: 3},
				{x: 1, y: 1},
				{x: 3, y: 0},
				{x: 1, y: -1},
				{x: 0, y: -3},
				{x: -1, y: -1},
				{x: -3, y: 0},
				{x: -1, y: 1},
			},
			want: &Polygon{corners: []Point{
				{x: 0, y: 3},
				{x: 1, y: 1},
				{x: 3, y: 0},
				{x: 1, y: -1},
				{x: 0, y: -3},
				{x: -1, y: -1},
				{x: -3, y: 0},
				{x: -1, y: 1},
			},
			},
			wantErr: false,
		},
		{
			name: "2 corner polygon: (0,3)-(1,1)",
			cors: []Point{
				{x: 0, y: 3},
				{x: 1, y: 1},
			},
			want:    &Polygon{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPolygon(tt.cors)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPolygon() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPolygon() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPolygon_Edges(t *testing.T) {

	tests := []struct {
		name string
		poly Polygon
		want []Line
	}{
		{
			name: "4 star polygon: (0,3)-(1,1)-(3,0)-(1,-1)-(0,-3)-(-1,-1)-(-3,0)-(-1,1)",
			poly: Polygon{corners: []Point{
				{x: 0, y: 3},
				{x: 1, y: 1},
				{x: 3, y: 0},
				{x: 1, y: -1},
				{x: 0, y: -3},
				{x: -1, y: -1},
				{x: -3, y: 0},
				{x: -1, y: 1},
			},
			},
			want: []Line{
				{startingPoint: Point{x: 0, y: 3}, endPoint: Point{x: 1, y: 1}},
				{startingPoint: Point{x: 1, y: 1}, endPoint: Point{x: 3, y: 0}},
				{startingPoint: Point{x: 3, y: 0}, endPoint: Point{x: 1, y: -1}},
				{startingPoint: Point{x: 1, y: -1}, endPoint: Point{x: 0, y: -3}},
				{startingPoint: Point{x: 0, y: -3}, endPoint: Point{x: -1, y: -1}},
				{startingPoint: Point{x: -1, y: -1}, endPoint: Point{x: -3, y: 0}},
				{startingPoint: Point{x: -3, y: 0}, endPoint: Point{x: -1, y: 1}},
				{startingPoint: Point{x: -1, y: 1}, endPoint: Point{x: 0, y: 3}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.poly.Edges(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Polygon.Edges() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPolygon_IsPointInside(t *testing.T) {
	fourStarPolygon := &Polygon{corners: []Point{
		{x: 0, y: 3},
		{x: 1, y: 1},
		{x: 3, y: 0},
		{x: 1, y: -1},
		{x: 0, y: -3},
		{x: -1, y: -1},
		{x: -3, y: 0},
		{x: -1, y: 1},
	}}
	doublePeakPolygon := &Polygon{corners: []Point{
		{x: -1, y: -1},
		{x: 1, y: 1},
		{x: 2, y: 0},
		{x: 3, y: 1},
		{x: 5, y: -1},
	}}
	z3_3586_Polygon := &Polygon{corners: []Point{
		{x: 437, y: 707},
		{x: 496, y: 704},
		{x: 495, y: 728},
		{x: 553, y: 733},
		{x: 555, y: 760},
		{x: 670, y: 760},
		{x: 674, y: 672},
		{x: 731, y: 672},
		{x: 733, y: 613},
		{x: 701, y: 610},
		{x: 702, y: 527},
		{x: 673, y: 524},
		{x: 675, y: 465},
		{x: 608, y: 462},
		{x: 609, y: 438},
		{x: 526, y: 438},
		{x: 518, y: 467},
		{x: 434, y: 470},
		{x: 432, y: 553},
		{x: 403, y: 556},
		{x: 402, y: 609},
		{x: 435, y: 615},
	}}

	tests := []struct {
		name string
		poly *Polygon
		p    Point
		want bool
	}{
		{
			name: "4 star polygon: (0,0) should be inside",
			poly: fourStarPolygon,
			p:    Point{x: 0, y: 0},
			want: true,
		},
		{
			name: "double peak polygon: (2,0.5) should be outside",
			poly: doublePeakPolygon,
			p:    Point{x: 2, y: 0.5},
			want: false,
		},
		{
			name: "double peak polygon: (2,-0.5) should be inside",
			poly: doublePeakPolygon,
			p:    Point{x: 2, y: -0.5},
			want: true,
		},
		{
			name: "double peak polygon: (1,0.5) should be inside",
			poly: doublePeakPolygon,
			p:    Point{x: 1, y: 0.5},
			want: true,
		},
		{
			name: "double peak polygon: (3,0.5) should be inside",
			poly: doublePeakPolygon,
			p:    Point{x: 3, y: 0.5},
			want: true,
		},
		{
			name: "double peak polygon: (-1,2) should be outside",
			poly: doublePeakPolygon,
			p:    Point{x: -1, y: 2},
			want: false,
		},
		{
			name: "double peak polygon: (-1,0.5) should be outside",
			poly: doublePeakPolygon,
			p:    Point{x: -1, y: 0.5},
			want: false,
		},
		{
			name: "double peak polygon: (4,0.5) should be outside",
			poly: doublePeakPolygon,
			p:    Point{x: 4, y: 0.5},
			want: false,
		},
		{
			name: "3586 z3 polygon: (472,719) should be outside",
			poly: z3_3586_Polygon,
			p:    Point{x: 472, y: 719},
			want: false,
		},
		{
			name: "3586 z3 polygon: (634,747) should be inside",
			poly: z3_3586_Polygon,
			p:    Point{x: 634, y: 747},
			want: true,
		},
		{
			name: "3586 z3 polygon: (382,616) should be outside",
			poly: z3_3586_Polygon,
			p:    Point{x: 382, y: 616},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.poly.IsPointInside(tt.p); got != tt.want {
				t.Errorf("Polygon.IsPointInside() = %v, want %v", got, tt.want)
			}
		})
	}
}

package geometry

import (
	"testing"
)

func TestCoordinate_String(t *testing.T) {
	tests := []struct {
		name string
		c    Coordinate
		want string
	}{
		{
			name: "test zero",
			c:    0.0,
			want: "0",
		},
		{
			name: "test 0.1",
			c:    0.1,
			want: "0.1",
		},
		{
			name: "test 0.10",
			c:    0.10,
			want: "0.1",
		},
		{
			name: "test -0.10",
			c:    -0.10,
			want: "-0.1",
		},
		{
			name: "test INF",
			c:    INF,
			want: "INF",
		},
		{
			name: "test NEG_INF",
			c:    NEG_INF,
			want: "-INF",
		},
		{
			name: "test MINUS_INF",
			c:    MINUS_INF,
			want: "-INF",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("Coordinate.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCoordinate_InRange(t *testing.T) {

	tests := []struct {
		name string
		c    Coordinate
		sc   Coordinate
		ec   Coordinate
		want bool
	}{
		{
			name: "test sc<ec inside",
			c:    0.0,
			sc:   -1,
			ec:   1,
			want: true,
		},
		{
			name: "test sc<ec outside low",
			c:    -2.0,
			sc:   -1,
			ec:   1,
			want: false,
		},
		{
			name: "test sc<ec outside high",
			c:    2.0,
			sc:   -1,
			ec:   1,
			want: false,
		},
		{
			name: "test sc<ec at sc",
			c:    -1.0,
			sc:   -1,
			ec:   1,
			want: true,
		},
		{
			name: "test sc<ec at ec",
			c:    1.0,
			sc:   -1,
			ec:   1,
			want: false,
		},
		{
			name: "test sc>ec inside",
			c:    0.0,
			sc:   1,
			ec:   -1,
			want: true,
		},
		{
			name: "test sc>ec outside low",
			c:    -2.0,
			sc:   1,
			ec:   -1,
			want: false,
		},
		{
			name: "test sc>ec outside high",
			c:    2.0,
			sc:   1,
			ec:   -1,
			want: false,
		},
		{
			name: "test sc>ec at sc",
			c:    1.0,
			sc:   1,
			ec:   -1,
			want: true,
		},
		{
			name: "test sc>ec at ec",
			c:    -1.0,
			sc:   1,
			ec:   -1,
			want: false,
		},
		{
			name: "test sc=ec inside",
			c:    1.0,
			sc:   1,
			ec:   1,
			want: true,
		},
		{
			name: "test sc=ec outside low",
			c:    -2.0,
			sc:   1,
			ec:   1,
			want: false,
		},
		{
			name: "test sc=ec outside high",
			c:    2.0,
			sc:   1,
			ec:   1,
			want: false,
		},
		{
			name: "test sc INF outside",
			c:    0,
			sc:   INF,
			ec:   1,
			want: false,
		},
		{
			name: "test sc INF inside",
			c:    2,
			sc:   INF,
			ec:   1,
			want: true,
		},
		{
			name: "test sc INF at ec",
			c:    1,
			sc:   INF,
			ec:   1,
			want: false,
		},
		{
			name: "test sc INF at INF",
			c:    INF,
			sc:   INF,
			ec:   1,
			want: true,
		},
		{
			name: "test ec INF inside",
			c:    2,
			sc:   1,
			ec:   INF,
			want: true,
		},
		{
			name: "test ec INF outside",
			c:    0,
			sc:   1,
			ec:   INF,
			want: false,
		},
		{
			name: "test ec INF at sc",
			c:    1,
			sc:   1,
			ec:   INF,
			want: true,
		},
		{
			name: "test ec INF at INF",
			c:    INF,
			sc:   1,
			ec:   INF,
			want: true,
		},
		{
			name: "test sc -INF inside",
			c:    -1,
			sc:   NEG_INF,
			ec:   0,
			want: true,
		},
		{
			name: "test sc -INF outside",
			c:    1,
			sc:   NEG_INF,
			ec:   0,
			want: false,
		},
		{
			name: "test sc -INF at -INF",
			c:    NEG_INF,
			sc:   NEG_INF,
			ec:   0,
			want: true,
		},
		{
			name: "test ec -INF inside",
			c:    -1,
			sc:   1,
			ec:   NEG_INF,
			want: true,
		},
		{
			name: "test ec -INF outside",
			c:    2,
			sc:   1,
			ec:   NEG_INF,
			want: false,
		},
		{
			name: "test ec -INF at sc",
			c:    1,
			sc:   1,
			ec:   NEG_INF,
			want: true,
		},
		{
			name: "test ec -INF at -INF",
			c:    NEG_INF,
			sc:   1,
			ec:   NEG_INF,
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.InRange(tt.sc, tt.ec); got != tt.want {
				t.Errorf("Coordinate.InRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

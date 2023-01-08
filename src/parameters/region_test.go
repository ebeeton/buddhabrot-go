// Package parameters defines parameters used to plot Buddhabrot images.
package parameters

import (
	"fmt"
	"math"
	"testing"
)

func TestPointInRegion(t *testing.T) {
	var tests = []struct {
		r    Region
		c    complex128
		want bool
	}{
		{Region{-2, 2, -2, 2}, complex(1, 1), true},
		{Region{-2, 2, -2, 2}, complex(2.1, 1), false},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%v,%t", tt.r, tt.c, tt.want)
		t.Run(testname, func(t *testing.T) {
			got := tt.r.pointInRegion(tt.c)

			if got != tt.want {
				t.Errorf("Got %t, want %t.", got, tt.want)
			}
		})
	}
}

func TestMatchAspectRatio(t *testing.T) {
	const (
		wantMinImag float64 = -0.92625
		wantMaxImag float64 = 0.92625
		delta               = 1e-10
	)
	r := Region{-2.0, 0.47, -1.12, 1.12}

	r.matchAspectRatio(1024, 768)

	if math.Abs(wantMinImag-r.MinImag) > delta {
		t.Errorf("Got %f, want %f.", r.MinImag, wantMinImag)
	} else if math.Abs(wantMaxImag-r.MaxImag) > delta {
		t.Errorf("Got %f, want %f.", r.MaxImag, wantMaxImag)
	}
}

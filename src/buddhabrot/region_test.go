// Package Buddhabrot plots images.
package buddhabrot

import (
	"fmt"
	"math"
	"testing"
)

func TestPointInRegion(t *testing.T) {
	var tests = []struct {
		r    region
		c    complex128
		want bool
	}{
		{region{-2, 2, -2, 2}, complex(1, 1), true},
		{region{-2, 2, -2, 2}, complex(2.1, 1), false},
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
	r := region{-2.0, 0.47, -1.12, 1.12}

	r.matchAspectRatio(1024, 768)

	if math.Abs(wantMinImag-r.minImag) > delta {
		t.Errorf("Got %f, want %f.", r.minImag, wantMinImag)
	} else if math.Abs(wantMaxImag-r.maxImag) > delta {
		t.Errorf("Got %f, want %f.", r.maxImag, wantMaxImag)
	}
}

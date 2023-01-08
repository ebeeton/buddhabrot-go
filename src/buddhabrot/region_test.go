// Package Buddhabrot plots images.
package buddhabrot

import (
	"fmt"
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

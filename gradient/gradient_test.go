// Package gradient defines functionality for generating linear color gradients.

package gradient

import (
	"testing"
)

func TestGetGradient(t *testing.T) {
	stops := []Stop{
		{Color: "#000000", Position: 0.0},
		{Color: "#FFFFFF", Position: 1.0},
	}

	wantStops := 1000
	g := GetGradient(stops, wantStops)

	if len(g) != wantStops {
		t.Errorf("Got stops %d, want %d.", len(g), wantStops)
	}
}

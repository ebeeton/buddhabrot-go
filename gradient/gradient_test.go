// Package gradient defines functionality for generating linear color gradients.

package gradient

import (
	"fmt"
	"image/color"
	"math"
	"testing"

	"github.com/lucasb-eyer/go-colorful"
)

func TestGetGradient(t *testing.T) {
	stops := []Stop{
		{Color: "#000000", Position: 0.0},
		{Color: "#FFFFFF", Position: 1.0},
	}

	wantStops := 3
	g := GetGradient(stops, wantStops)

	black := color.RGBA{R: 0, G: 0, B: 0, A: math.MaxUint8}
	grey := color.RGBA{R: math.MaxUint8 >> 1, G: math.MaxUint8 >> 1, B: math.MaxUint8 >> 1, A: math.MaxUint8}
	white := color.RGBA{R: math.MaxUint8, G: math.MaxUint8, B: math.MaxUint8, A: math.MaxUint8}

	if len(g) != wantStops {
		t.Errorf("Got stops %d, want %d.", len(g), wantStops)
	} else if g[0] != black {
		t.Errorf("Got black %v, want %v.", g[0], black)
	} else if g[1] != grey {
		t.Errorf("Got grey %v, want %v.", g[1], grey)
	} else if g[2] != white {
		t.Errorf("Got white %v, want %v.", g[2], white)
	}
}

func TestColorfulToColor(t *testing.T) {
	var tests = []struct {
		c    colorful.Color
		want color.RGBA
	}{
		{colorful.Color{R: 1, G: 1, B: 1}, color.RGBA{R: 255, G: 255, B: 255, A: 255}},
		{colorful.Color{R: 1, G: 0, B: 0}, color.RGBA{R: 255, G: 0, B: 0, A: 255}},
		{colorful.Color{R: 0, G: 1, B: 0}, color.RGBA{R: 0, G: 255, B: 0, A: 255}},
		{colorful.Color{R: 0, G: 0, B: 1}, color.RGBA{R: 0, G: 0, B: 255, A: 255}},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v", tt.c)
		t.Run(testname, func(t *testing.T) {
			got := colorfulToColor(tt.c)
			if got != tt.want {
				t.Errorf("Got %v, want %v.", got, tt.want)
			}
		})
	}
}

func TestGetInterpolatedColor(t *testing.T) {
	table := gradientTable{
		stop{col: colorful.Color{R: 0, G: 0, B: 0}, pos: 0},
		stop{col: colorful.Color{R: 1, G: 1, B: 1}, pos: 1},
	}
	want := colorful.Color{R: 0.5, G: 0.5, B: 0.5}

	got := table.getInterpolatedColor(0.5)

	if got != want {
		t.Errorf("Got %v, want %v.", got, want)
	}
}

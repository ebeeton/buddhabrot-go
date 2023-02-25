// Package gradient defines functionality for generating linear color gradients.

package gradient

import (
	"fmt"
	"image/color"
	"testing"
)

func TestColorFromHex(t *testing.T) {
	hexColor := "#eb4034"
	want := color.RGBA{R: 235, G: 64, B: 52, A: 255}

	got, err := colorFromHex(hexColor)
	if err != nil {
		t.Error(err.Error())
	} else if want != got {
		t.Errorf("Got %v, want %v.", got, want)
	}
}

func TestColorFromHexError(t *testing.T) {
	hexColor := "#eb403"

	_, err := colorFromHex(hexColor)
	if err == nil {
		t.Error("Want error, got nil.")
	}
}

func TestLerp(t *testing.T) {
	var tests = []struct {
		first, second uint8
		stop          float64
		want          uint8
	}{
		{64, 192, 0.5, 128},
		{0, 255, 0.25, 63},
		{0, 255, 0.75, 191},
		{0, 255, 0, 0},
		{0, 255, 1.0, 255},
		{255, 255, 0.5, 255},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%d,%f", tt.first, tt.second, tt.stop)
		t.Run(testname, func(t *testing.T) {
			got := lerp(tt.first, tt.second, tt.stop)
			if got != tt.want {
				t.Errorf("Got %d, want %d.", got, tt.want)
			}
		})
	}
}

// Package gradient defines functionality for generating linear color gradients.

package gradient

import (
	"image/color"
	"testing"
)

func TestColorFromHex(t *testing.T) {
	hexColor := "#eb4034"
	want := color.RGBA{R: 235, G: 64, B: 52}

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

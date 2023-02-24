// Package gradient defines functionality for generating linear color gradients.

package gradient

import (
	"encoding/hex"
	"errors"
	"image/color"
	"strings"
)

func colorFromHex(h string) (color.RGBA, error) {
	data, err := hex.DecodeString(strings.Trim(h, "#"))
	if err != nil {
		return color.RGBA{}, err
	} else if len(data) != 3 {
		return color.RGBA{}, errors.New("argument must be six hex characters")
	}
	return color.RGBA{
		R: data[0],
		G: data[1],
		B: data[2],
	}, nil
}

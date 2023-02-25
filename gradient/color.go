// Package gradient defines functionality for generating linear color gradients.

package gradient

import (
	"encoding/hex"
	"errors"
	"image/color"
	"math"
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
		A: math.MaxUint8,
	}, nil
}

func linearInterpolate(first, second color.RGBA, pos float64) color.RGBA {
	return color.RGBA{
		R: lerp(first.R, second.R, pos),
		G: lerp(first.G, second.G, pos),
		B: lerp(first.B, second.B, pos),
		A: lerp(first.A, second.A, pos),
	}
}

func lerp(first, second uint8, stop float64) uint8 {
	return uint8(float64(first)*(1.0-stop) + float64(second)*stop)
}

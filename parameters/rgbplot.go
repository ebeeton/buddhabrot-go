// Package parameters defines parameters used to plot Buddhabrot images.
package parameters

import (
	"errors"
)

const (
	Red = iota
	Green
	Blue
)

// RgbPlot represents the parameters to plot an RGB Buddhabrot image of a given
// width and height. Each color channel is configured independently. Region
// defines the region on the complex plane to plot.
type RgbPlot struct {
	Red,
	Green,
	Blue Channel
	Region Region
	Width,
	Height int
}

// GetChannel returns the red, green, or blue Channel by passing in the Red,
// Green, or Blue constants from this package.
func (p RgbPlot) GetChannel(index int) (Channel, error) {
	switch index {
	case Red:
		return p.Red, nil
	case Green:
		return p.Green, nil
	case Blue:
		return p.Blue, nil
	default:
		return Channel{}, errors.New("invalid index")
	}
}

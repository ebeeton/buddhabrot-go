// Package Buddhabrot plots images.
package buddhabrot

import (
	"math"

	"github.com/ebeeton/buddhalbrot-go/parameters"
)

const (
	Channels = 3
)

func Plot(plot parameters.RgbPlot) [][]uint32 {
	counter := make([][]uint32, Channels)
	for i := range counter {
		counter[i] = make([]uint32, plot.Height*plot.Width)
	}
	return counter
}

func isInMandelbrotSet(c complex128, maxIterations int) (bool, int) {
	const bailout float64 = 2
	z := c
	for i := 0; i < maxIterations; i++ {
		if math.Abs(real(z)) > bailout || math.Abs(imag(z)) > bailout {
			return false, i
		}
		z = z*z + c
	}

	return true, maxIterations
}

func linearScale(val, minScaleFrom, maxScaleFrom, minScaleTo, maxScaleTo float64) float64 {
	return (val-minScaleFrom)/(maxScaleFrom-minScaleFrom)*(maxScaleTo-minScaleTo) + minScaleTo
}

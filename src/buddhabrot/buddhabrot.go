// Package Buddhabrot plots images.
package buddhabrot

import (
	"math"

	"github.com/ebeeton/buddhalbrot-go/parameters"
)

const (
	Channels                = 3
	ComplexPlaneMin float64 = -2
	ComplexPlaneMax float64 = 2
)

func Plot(plot parameters.RgbPlot) [][]uint32 {
	counter := make([][]uint32, Channels)
	for i := range counter {
		counter[i] = make([]uint32, plot.Height*plot.Width)
	}

	plotChannel(plot.Red)
	plotChannel(plot.Green)
	plotChannel(plot.Blue)

	return counter
}

func plotChannel(c parameters.Channel) {

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

func randomPointOnComplexPlane() complex128 {
	return complex128(0)
}

// Package Buddhabrot plots images.
package buddhabrot

import (
	"math"
	"math/rand"

	"github.com/ebeeton/buddhalbrot-go/parameters"
)

const (
	bailout         float64 = 2
	bailoutTimesTwo float64 = bailout * 2
	channels                = 3
	complexPlaneMin float64 = -2
	complexPlaneMax float64 = 2
)

// Plot plots points not in the Mandelbrot set as they escape to infinity. It
// increments a counter for each point on the plane every time an orbit passes
// through it. The counter is returned as a slice of three channels
// corresponding to red, green, and blue in an RGB image. Each channel is a
// slice of uint32 the length of the image width times height.
func Plot(plot parameters.RgbPlot) [][]uint32 {
	counter := make([][]uint32, channels)
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

func isInMandelbrotSet(c complex128, maxIterations int) bool {
	z := c
	for i := 0; i < maxIterations; i++ {
		if math.Abs(real(z)) > bailout || math.Abs(imag(z)) > bailout {
			return false
		}
		z = z*z + c
	}

	return true
}

func linearScale(val, minScaleFrom, maxScaleFrom, minScaleTo, maxScaleTo float64) float64 {
	return (val-minScaleFrom)/(maxScaleFrom-minScaleFrom)*(maxScaleTo-minScaleTo) + minScaleTo
}

func randomPointOnComplexPlane() complex128 {
	// Generate r and imaginary values from -2 to 2.
	r := rand.Float64()*bailoutTimesTwo - bailout
	i := rand.Float64()*bailoutTimesTwo - bailout
	return complex(r, i)
}

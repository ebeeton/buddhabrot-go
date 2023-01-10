// Package Buddhabrot plots images.
package buddhabrot

import (
	"log"
	"math"
	"math/rand"
	"sync/atomic"

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

	plotChannel(parameters.Red, counter[parameters.Red], plot)
	plotChannel(parameters.Green, counter[parameters.Green], plot)
	plotChannel(parameters.Blue, counter[parameters.Blue], plot)

	return counter
}

func plotChannel(channelIndex int, counter []uint32, plot parameters.RgbPlot) {
	channel, err := plot.GetChannel(channelIndex)
	if err != nil {
		log.Printf("cannot plot channel %d", channelIndex)
		return
	}

	for i := 0; i < channel.SampleSize; i++ {
		point := randomPointNotInMandelbrotSet(channel.MaxSampleIterations)
		orbits := plotOrbits(point, channel.MaxIterations, plot.Region)

		for _, v := range orbits {
			// Convert from complex to image space.
			pX := int(linearScale(real(v), plot.Region.MinReal, plot.Region.MaxReal, 0, float64(plot.Width)))
			pY := int(linearScale(imag(v), plot.Region.MinImag, plot.Region.MaxImag, 0, float64(plot.Height)))
			index := pY*plot.Width + pX

			// The same counter could be incremented when run concurrently,
			// so increment as an atomic operation.
			atomic.AddUint32(&counter[index], 1)
		}
	}
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

func randomPointNotInMandelbrotSet(maxIterations int) complex128 {
	for {
		// Generate r and imaginary values from -2 to 2.
		r := rand.Float64()*bailoutTimesTwo - bailout
		i := rand.Float64()*bailoutTimesTwo - bailout
		p := complex(r, i)
		if !isInMandelbrotSet(p, maxIterations) {
			return p
		}
	}
}

func plotOrbits(c complex128, maxIterations int, r parameters.Region) []complex128 {
	var orbits []complex128
	z := c
	for i := 0; i < maxIterations; i++ {
		if math.Abs(real(z)) > bailout || math.Abs(imag(z)) > bailout {
			// Point has escaped to infinity.
			return orbits
		} else if r.PointInRegion(z) {
			// Only save orbits within the plot region.
			orbits = append(orbits, z)
		}
		z = z*z + c
	}
	return orbits
}

// Package Buddhabrot plots images.
package buddhabrot

import (
	"image"
	"image/color"
	"log"
	"math"
	"math/rand"
	"sync"
	"sync/atomic"

	"github.com/ebeeton/buddhalbrot-go/parameters"
	"github.com/ebeeton/buddhalbrot-go/timer"
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
func Plot(plot parameters.RgbPlot) (*image.RGBA, error) {
	defer timer.Timer("Plot")()
	log.Printf("Plot started with params: %+v.", plot)
	counter := make([][]uint32, channels)
	var channelMax [channels]uint32
	for i := range counter {
		counter[i] = make([]uint32, plot.Height*plot.Width)
		max, err := plotChannel(i, counter[i], plot)
		if err != nil {
			return nil, err
		}
		channelMax[i] = max
	}
	log.Println("Highest count per channel:", channelMax)

	// Assign each pixel color channel a value based on how many times an orbit
	// "passed through" it.
	img := image.NewRGBA(image.Rect(0, 0, plot.Width, plot.Height))
	pixelStride := img.Stride >> 2
	for y := 0; y < plot.Height; y++ {
		for x := 0; x < plot.Width; x++ {
			idx := pixelStride*y + x
			img.SetRGBA(x, y, color.RGBA{
				R: uint8(float64(counter[parameters.Red][idx]) / float64(channelMax[parameters.Red]) * math.MaxUint8),
				G: uint8(float64(counter[parameters.Green][idx]) / float64(channelMax[parameters.Green]) * math.MaxUint8),
				B: uint8(float64(counter[parameters.Blue][idx]) / float64(channelMax[parameters.Blue]) * math.MaxUint8),
				A: math.MaxUint8,
			})
		}
	}
	log.Println("Plot complete.")
	return img, nil
}

func plotChannel(channelIndex int, counter []uint32, plot parameters.RgbPlot) (uint32, error) {
	channel, err := plot.GetChannel(channelIndex)
	if err != nil {
		return 0, err
	}

	log.Printf("Channel %d plot started.", channelIndex)
	max := uint32(0)
	var wg sync.WaitGroup
	for i := 0; i < channel.SampleSize; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			point := randomPointNotInMandelbrotSet(channel.MaxIterations)
			orbits := plotOrbits(point, channel.MaxIterations, plot.Region)

			for _, v := range orbits {
				// Convert from complex to image space.
				pX := int(linearScale(real(v), plot.Region.MinReal, plot.Region.MaxReal, 0, float64(plot.Width)))
				pY := int(linearScale(imag(v), plot.Region.MinImag, plot.Region.MaxImag, 0, float64(plot.Height)))
				index := pY*plot.Width + pX

				// The same counter could be incremented by more than one
				// goroutine so increment as an atomic operation.
				if val := atomic.AddUint32(&counter[index], 1); val > max {
					max = val
				}
			}
		}()
	}
	log.Printf("Channel %d plot complete.", channelIndex)
	return max, nil
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

func lerp(first, second uint8, stop float64) uint8 {
	return uint8(float64(first)*(1.0-stop) + float64(second)*stop)
}

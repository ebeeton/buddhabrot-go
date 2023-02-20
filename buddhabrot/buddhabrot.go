// Package Buddhabrot plots images.
package buddhabrot

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"log"
	"math"
	"math/rand"
	"os"
	"sort"
	"sync"
	"sync/atomic"

	"github.com/ebeeton/buddhalbrot-go/gradient"
	"github.com/ebeeton/buddhalbrot-go/parameters"
	"github.com/ebeeton/buddhalbrot-go/timer"
)

const (
	bailout         float64 = 2
	bailoutTimesTwo float64 = bailout * 2
	channels                = 3
	complexPlaneMin float64 = -2
	complexPlaneMax float64 = 2
	paletteColors           = 256
)

// Plot iteratively plots points not in the Mandelbrot set as they escape to
// infinity. It increments a counter for each point on the complex plane every
// time an orbit passes through it. The counter value is used to determine the
// grayscale value of each pixel in the returned image.RGBA pointer.
func Plot(plot parameters.Plot) (*image.RGBA, error) {
	defer timer.Timer("Plot")()
	log.Printf("Plot started with params: %+v.", plot)
	counter := make([]uint32, plot.Height*plot.Width)
	var wg sync.WaitGroup
	for i := 0; i < plot.SampleSize; i++ {
		wg.Add(1)
		go plotPoint(plot, counter, &wg)
	}
	wg.Wait()

	if plot.DumpCounterFile {
		dumpCounterFile(counter)
	}

	// Get the gradient palette used to color pixels based on hit count.
	g := getPaletteMap(counter, plot.Gradient)

	// Assign each pixel a color based on how many times an orbit "passed
	// through" it.
	img := image.NewRGBA(image.Rect(0, 0, plot.Width, plot.Height))
	pixelStride := img.Stride >> 2
	for y := 0; y < plot.Height; y++ {
		for x := 0; x < plot.Width; x++ {
			cIdx := pixelStride*y + x
			img.SetRGBA(x, y, g[counter[cIdx]])
		}
	}
	return img, nil
}

func getPaletteMap(counter []uint32, stops []gradient.Stop) map[uint32]color.RGBA {
	defer timer.Timer("getPaletteMap")()
	// Get unique orbit counts, and assign each one a color based on where it
	// would fall in the desired gradient.
	paletteMap := make(map[uint32]color.RGBA)
	for _, c := range counter {
		paletteMap[c] = color.RGBA{}
	}
	log.Println("Unique orbit counts:", len(paletteMap))

	// Extract the unique orbit counts.
	keys := make([]uint32, len(paletteMap))
	i := 0
	for k := range paletteMap {
		keys[i] = k
		i++
	}

	// Sort the unique orbit counts.
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	// Map each unique orbit count to a color.
	g := gradient.GetGradient(stops, len(keys))
	for i, k := range keys {
		paletteMap[k] = g[i]
	}

	return paletteMap
}

func plotPoint(plot parameters.Plot, counter []uint32, wg *sync.WaitGroup) {
	defer wg.Done()
	point := randomPointNotInMandelbrotSet(plot.MaxIterations)
	orbits := plotOrbits(point, plot.MaxIterations, plot.Region)

	for _, v := range orbits {
		// Convert from complex to image space.
		pX := int(linearScale(real(v), plot.Region.MinReal, plot.Region.MaxReal, 0, float64(plot.Width)))
		pY := int(linearScale(imag(v), plot.Region.MinImag, plot.Region.MaxImag, 0, float64(plot.Height)))
		index := pY*plot.Width + pX

		// The same counter could be incremented by more than one
		// goroutine so increment as an atomic operation.
		atomic.AddUint32(&counter[index], 1)
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

func lerp(first, second uint8, stop float64) uint8 {
	return uint8(float64(first)*(1.0-stop) + float64(second)*stop)
}

func dumpCounterFile(counter []uint32) {
	if err := os.MkdirAll("log", os.ModePerm); err != nil {
		log.Printf("Failed to create log director: %v", err.Error())
		return
	}
	f, err := os.Create("log/counter.txt")
	if err != nil {
		log.Printf("Failed to create dump file: %v", err.Error())
		return
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	for _, c := range counter {
		if _, err := w.WriteString(fmt.Sprintf("%d\n", c)); err != nil {
			log.Printf("WriteString error: %v", err.Error())
			break
		}
	}
	w.Flush()
}

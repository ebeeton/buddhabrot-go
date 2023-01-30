// Package parameters defines parameters used to plot Buddhabrot images.
package parameters

import "github.com/ebeeton/buddhalbrot-go/gradient"

// Plot represents the parameters to plot a Buddhabrot image of a given width
// and height. SampleSize determines the number of sample points used.
// MaxIterations is the maximum number of iterations to determine if a sample
// point is in the Mandelbrot set. It is also the maximum number of iterations
// to count orbits that pass through points visible in the plot. Region defines
// the region on the complex plane to plot. Gradient is a slice of colors and
// their positions in a linear gradient used to color the image. If set to
// true, the `dumpCounterFile` property wil dump the orbit counts per pixel to a
// file called counter.txt in the log directory.
type Plot struct {
	SampleSize,
	MaxIterations int `validate:"gte=1"`
	Region Region
	Width,
	Height int `validate:"gte=1"`
	Gradient        []gradient.Stop `validate:"validateGradient"`
	DumpCounterFile bool
}

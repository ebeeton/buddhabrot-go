// Package parameters defines parameters used to plot Buddhabrot images.
package parameters

// Plot represents the parameters to plot a Buddhabrot image of a given width
// and height. SampleSize determines the number of sample points used.
// MaxIterations is the maximum number of iterations to determine if a sample
// point is in the Mandelbrot set. It is also the maximum number of iterations
// to count orbits that pass through points visible in the plot. Region defines
//the region on the complex plane to plot.
type Plot struct {
	SampleSize,
	MaxIterations int
	Region Region
	Width,
	Height int
}

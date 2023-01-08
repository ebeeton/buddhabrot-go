// Package parameters defines parameters used to plot Buddhabrot images.
package parameters

// Channel represents the parameters for a single channel in an RGB image.
// SampleSize determines the number of sample points used. MaxSampleIterations
// is the maximum number of iterations to determine if a sample point is in the
// Mandelbrot set. MaxIterations is the maximum number of iterations to count
// orbits that pass through points visible in the plot.
type Channel struct {
	SampleSize,
	MaxSampleIterations,
	MaxIterations uint64
}

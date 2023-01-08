// Package Buddhabrot plots images.
package buddhabrot

// Region is a region on the complex plane defined by minimum and maximum real
// and imaginary parts.
type region struct {
	minReal,
	maxReal,
	minImag,
	maxImag float64
}

func (r region) pointInRegion(c complex128) bool {
	return real(c) >= r.minReal &&
		real(c) <= r.maxReal &&
		imag(c) >= r.minImag &&
		imag(c) <= r.maxImag
}

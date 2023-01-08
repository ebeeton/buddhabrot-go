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

func (r *region) matchAspectRatio(width, height int) {
	complexWidth := r.maxReal - r.minReal
	aspectRatio := float64(height) / float64(width)
	newComplexHeight := complexWidth * aspectRatio
	var halfComplexHeightDelta = (newComplexHeight - (r.maxImag - r.minImag)) / 2.0
	r.minImag -= halfComplexHeightDelta
	r.maxImag += halfComplexHeightDelta
}

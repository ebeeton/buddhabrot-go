// Package parameters defines parameters used to plot Buddhabrot images.
package parameters

// Region is a region on the complex plane defined by minimum and maximum real
// and imaginary parts.
type Region struct {
	MinReal,
	MaxReal,
	MinImag,
	MaxImag float64
}

func (r Region) pointInRegion(c complex128) bool {
	return real(c) >= r.MinReal &&
		real(c) <= r.MaxReal &&
		imag(c) >= r.MinImag &&
		imag(c) <= r.MaxImag
}

func (r *Region) matchAspectRatio(width, height int) {
	complexWidth := r.MaxReal - r.MinReal
	aspectRatio := float64(height) / float64(width)
	newComplexHeight := complexWidth * aspectRatio
	var halfComplexHeightDelta = (newComplexHeight - (r.MaxImag - r.MinImag)) / 2.0
	r.MinImag -= halfComplexHeightDelta
	r.MaxImag += halfComplexHeightDelta
}

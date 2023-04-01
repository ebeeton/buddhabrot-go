// Package parameters defines parameters used to plot Buddhabrot images.

package parameters

// Stop is a hexadecimal color value and its position in a linear gradient from
// 0 to 1.
type Stop struct {
	Color    string
	Position float64
}

// Package parameters defines parameters used to plot Buddhabrot images.
package parameters

// RgbPlot represents the parameters to plot an RGB Buddhabrot image of a given
// width and height. Each color channel is configured independently.
type RgbPlot struct {
	Red,
	Green,
	Blue Channel
	Width,
	Height int
}

// Package parameters defines parameters used to plot Buddhabrot images.
package parameters

// RgbPlot represents the parameters to plot an RGB Buddhabrot image of a given
// width and height. Each color channel is configured independently. Region
// defines the region on the complex plane to plot.
type RgbPlot struct {
	Red,
	Green,
	Blue Channel
	Region Region
	Width,
	Height int
}

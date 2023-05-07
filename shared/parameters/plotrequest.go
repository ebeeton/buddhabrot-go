// Package parameters defines parameters used to plot Buddhabrot images.
package parameters

// Type PlotRequest represents a request for a Buddhabrot image. It contains the
// database ID for the result, and the parameters used to perform the plot.
type PlotRequest struct {
	Id   uint
	Plot Plot
}

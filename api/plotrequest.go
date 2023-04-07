package main

import "github.com/ebeeton/buddhabrot-go/parameters"

// Type PlotRequest represents a request for a Buddhabrot image. It contains the
// database ID for the result, and the parameters used to perform the plot.
type PlotRequest struct {
	Id   int64
	Plot parameters.Plot
}

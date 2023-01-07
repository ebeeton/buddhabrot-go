// Package Buddhabrot plots images.
package buddhabrot

import "github.com/ebeeton/buddhalbrot-go/parameters"

const (
	Channels = 3
)

func Plot(plot parameters.RgbPlot) [][]uint32 {
	counter := make([][]uint32, Channels)
	for i := range counter {
		counter[i] = make([]uint32, plot.Height*plot.Width)
	}
	return counter
}

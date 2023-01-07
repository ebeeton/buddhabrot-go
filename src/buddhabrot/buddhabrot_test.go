// Package Buddhabrot plots images.
package buddhabrot

import (
	"testing"

	"github.com/ebeeton/buddhalbrot-go/parameters"
)

func TestPlot(t *testing.T) {
	plot := parameters.RgbPlot{
		Red: parameters.Channel{
			SampleSize:          100000000,
			MaxSampleIterations: 1000,
			MaxIterations:       1000,
		},
		Green: parameters.Channel{
			SampleSize:          100000000,
			MaxSampleIterations: 1000,
			MaxIterations:       1000,
		},
		Blue: parameters.Channel{
			SampleSize:          100000000,
			MaxSampleIterations: 1000,
			MaxIterations:       1000,
		},
		Width:  1024,
		Height: 768,
	}

	result := Plot(plot)

	if len(result) != Channels {
		t.Errorf("Expected %d channels, got %d.", Channels, len(result))
	} else if len(result[0]) != int(plot.Width)*int(plot.Height) {
		t.Errorf("Expected channel length %d, got %d",
			int(plot.Width)*int(plot.Height), len(result[0]))
	}

}

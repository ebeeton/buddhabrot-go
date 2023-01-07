// Package Buddhabrot plots images.
package buddhabrot

import (
	"fmt"
	"math"
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

func TestLinearScale(t *testing.T) {
	// "Close enough" from https://stackoverflow.com/a/47969546
	const delta = 1e-10
	var tests = []struct {
		val, minScaleFrom, maxScaleFrom, minScaleTo, maxScaleTo, want float64
	}{
		{5, 0, 10, 0, 100, 50},
		{50, 0, 100, 0, 1, 0.5},
		{75, 0, 100, 0, 1, 0.75},
		{0, -2, 0.47, 0, 3072, 2487.449392712551},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%f,%f,%f,%f,%f", tt.val, tt.minScaleFrom, tt.maxScaleFrom, tt.minScaleTo, tt.maxScaleTo)
		t.Run(testname, func(t *testing.T) {
			result := linearScale(tt.val, tt.minScaleFrom, tt.maxScaleFrom, tt.minScaleTo, tt.maxScaleTo)
			if math.Abs(tt.want-result) > delta {
				t.Errorf("Got %f, want %f", result, tt.want)
			}
		})
	}
}

func TestIsInMandelbrotSet(t *testing.T) {
	var tests = []struct {
		c             complex128
		maxIterations int
		isInSet       bool
		iterations    int
	}{
		{complex(0, 0), 100, true, 100},
		{complex(1, 1), 100, false, 1},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%d,%t,%d", tt.c, tt.maxIterations, tt.isInSet, tt.iterations)
		t.Run(testname, func(t *testing.T) {
			isInSet, iterations := isInMandelbrotSet(tt.c, tt.maxIterations)

			if isInSet != tt.isInSet {
				t.Errorf("Got %t, want %t", isInSet, tt.isInSet)
			} else if iterations != tt.iterations {
				t.Errorf("Got %d, want %d", iterations, tt.iterations)
			}
		})
	}
}

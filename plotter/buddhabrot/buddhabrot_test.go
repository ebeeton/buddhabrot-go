// Package Buddhabrot plots images.
package buddhabrot

import (
	"fmt"
	"math"
	"testing"

	"github.com/ebeeton/buddhabrot-go/plotter/parameters"
)

func TestPlot(t *testing.T) {
	plot := parameters.Plot{
		SampleSize:    100,
		MaxIterations: 10,
		Width:         256,
		Height:        128,
		Gradient: []parameters.Stop{
			{Color: "#000000"},
			{Color: "#FFFFFF"},
		},
	}

	got := Plot(plot)

	if got.Rect.Dx() != plot.Width {
		t.Errorf("Got width %d, want %d", got.Rect.Dx(), plot.Width)
	} else if got.Rect.Dy() != plot.Height {
		t.Errorf("Got width %d, want %d", got.Rect.Dy(), plot.Height)
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
	}{
		{complex(0, 0), 100, true},
		{complex(1, 1), 100, false},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v,%d,%t", tt.c, tt.maxIterations, tt.isInSet)
		t.Run(testname, func(t *testing.T) {
			isInSet := isInMandelbrotSet(tt.c, tt.maxIterations)

			if isInSet != tt.isInSet {
				t.Errorf("Got %t, want %t", isInSet, tt.isInSet)
			}
		})
	}
}

func TestRandomPointNotInMandelbrotSet(t *testing.T) {
	point := randomPointNotInMandelbrotSet(100)
	if real(point) > complexPlaneMax || real(point) < complexPlaneMin {
		t.Errorf("Got real %f, want between %f and %f.", real(point),
			complexPlaneMin, complexPlaneMax)
	} else if imag(point) > complexPlaneMax || imag(point) < complexPlaneMin {
		t.Errorf("Got imaginary %f, want between %f and %f.", imag(point),
			complexPlaneMin, complexPlaneMax)
	}
}

func TestPlotOrbits(t *testing.T) {
	c := complex(0.42, 0.42)
	r := parameters.Region{MinReal: -2, MaxReal: 2, MinImag: -2, MaxImag: 2}

	got := plotOrbits(c, 10, r)

	want := []complex128{(0.42 + 0.42i),
		(0.42 + 0.7727999999999999i),
		(-0.0008198399999999606 + 1.0691519999999999i),
		(-0.7230853269663742 + 0.41824693284864006i),
		(0.7679218932367734 + -0.18485644038308408i),
		(0.9755321305612457 + 0.13608938464802267i),
		(1.3531426171434857 + 0.6855191347049088i)}

	if len(got) != len(want) {
		t.Errorf("Got %d orbits, want %d.", len(got), len(want))
	}

	for i, v := range want {
		if got[i] != v {
			t.Errorf("Index %d got %v, want %v.", i, got[i], v)
		}
	}
}

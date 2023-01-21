// Package histogram equalizes the distribution of orbit counts within a plot.

package histogram

import "testing"

func TestNormalize(t *testing.T) {
	h := Histogram{
		117: 2,
		118: 1,
		119: 2,
		120: 3,
		121: 1,
		122: 6,
		123: 4,
		124: 10,
	}

	h.Normalize(29)

	want := 1.0
	got := 0.0
	for _, v := range h {
		got += v
	}

	if got != want {
		t.Errorf("Got %f, want %f.", got, want)
	}
}

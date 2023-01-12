// Package parameters defines parameters used to plot Buddhabrot images.
package parameters

import (
	"fmt"
	"testing"
)

func TestGetChannel(t *testing.T) {
	p := RgbPlot{Red: Channel{SampleSize: 1},
		Green: Channel{SampleSize: 2},
		Blue:  Channel{SampleSize: 3}}

	var tests = []struct {
		index   int
		channel Channel
		err     bool
	}{
		{Red, p.Red, false},
		{Green, p.Green, false},
		{Blue, p.Blue, false},
		{5, Channel{}, true},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%d,%v", tt.index, tt.err)
		t.Run(testname, func(t *testing.T) {
			c, err := p.GetChannel(tt.index)
			if c != tt.channel {
				t.Errorf("Got %v, want %v.", c, tt.channel)
			} else if tt.err && err == nil {
				t.Errorf("Got nil, want error.")
			} else if !tt.err && err != nil {
				t.Errorf("Got an error, want nil.")
			}
		})
	}
}

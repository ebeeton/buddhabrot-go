// Package gradient defines functionality for generating linear color gradients.

package gradient

import (
	"image/color"
	"math"

	"github.com/lucasb-eyer/go-colorful"
)

// This code used the gradient example from Lucas Beyer's excellent
// go-colorful library.
type stop struct {
	col colorful.Color
	pos float64
}

type gradientTable []stop

func GetGradient(stops []Stop, count int) []color.RGBA {
	t := gradientTable{}
	p := []color.RGBA{}
	for _, s := range stops {
		c, err := colorful.Hex(s.Color)
		if err != nil {
			panic("GetGradient: " + err.Error())
		}
		t = append(t, stop{col: c, pos: s.Position})
	}

	for i := 0; i < count; i++ {
		// TODO:: fix an off-by-one error here that prevents the stop at 1.0
		// from being the last color in the resulting slice.
		col := t.getInterpolatedColor(float64(i) / float64(count))
		p = append(p, color.RGBA{
			R: uint8(math.MaxUint8 * col.R),
			G: uint8(math.MaxUint8 * col.G),
			B: uint8(math.MaxUint8 * col.B),
			A: math.MaxUint8,
		})
	}
	return p
}

func (gt gradientTable) getInterpolatedColor(t float64) colorful.Color {
	for i := 0; i < len(gt)-1; i++ {
		c1 := gt[i]
		c2 := gt[i+1]
		if c1.pos <= t && t <= c2.pos {
			// Blend the two points we're between.
			t := (t - c1.pos) / (c2.pos - c1.pos)
			return c1.col.BlendHcl(c2.col, t).Clamped()
		}
	}

	// We're not between any points, so return the last color.
	return gt[len(gt)-1].col
}

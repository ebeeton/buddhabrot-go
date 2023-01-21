// Package histogram equalizes the distribution of orbit counts within a plot.

package histogram

type Histogram map[uint32]float64

// Normalize normalizes the orbit counts for each pixel so that they sum to 1.
func (h *Histogram) Normalize(pixelCount int) {
	for k, v := range *h {
		(*h)[k] = v / float64(pixelCount)
	}
}

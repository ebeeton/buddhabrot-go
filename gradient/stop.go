// Package gradient defines functionality for generating linear color gradients.

package gradient

// Stop is a hexadecimal color value and its position in a linear gradient from
// 0 to 1.
type Stop struct {
	Color    string  `validate:"hexcolor"`
	Position float64 `validate:"gte=0,lte=1"`
}

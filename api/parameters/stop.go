// Package parameters defines parameters used to plot Buddhabrot images.

package parameters

import "github.com/go-playground/validator/v10"

// Stop is a hexadecimal color value and its position in a linear gradient from
// 0 to 1.
type Stop struct {
	Color    string  `validate:"hexcolor"`
	Position float64 `validate:"gte=0,lte=1"`
}

// ValidateStops validates that the state of a slice of Stops.
func ValidateStops(fl validator.FieldLevel) bool {
	// TODO:: How do you add specific error messages?
	stops := fl.Field().Interface().([]Stop)
	if len(stops) < 2 {
		return false
	} else if stops[0].Position != 0 {
		return false
	} else if stops[len(stops)-1].Position != 1 {
		return false
	}

	return true
}

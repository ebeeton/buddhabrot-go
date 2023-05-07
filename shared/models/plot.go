// Package models defines shared database models.
package models

import "gorm.io/gorm"

// A Buddhabrot plot represented as the parameteres used to generate it, and the
// filename of the image once the plot has completed.
type Plot struct {
	gorm.Model
	Params   string `gorm:"type:json not null"`
	Filename string `gorm:"type:varchar(255) null"`
}

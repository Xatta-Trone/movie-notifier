package entities

import "gorm.io/gorm"

type Tracker struct {
	gorm.Model
	Keyword  string
	IsParsed bool `gorm:"default:0"`
}

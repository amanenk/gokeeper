package table

import "gorm.io/gorm"

type Table struct {
	gorm.Model
	Number int  `gorm:"unique" json:"number"`
	Seats  int  `json:"seats"`
	IsBusy bool `json:"isBusy" validate:"required"`
}

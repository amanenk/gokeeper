package meal

import "gorm.io/gorm"

type Meal struct {
	gorm.Model
	Name  string `json:"name"`
	Price uint   `json:"price"` //price in cents
}

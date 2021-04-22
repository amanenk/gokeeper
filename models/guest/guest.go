package guest

import (
	"github.com/fdistorted/gokeeper/models/ordered-meal"
	"gorm.io/gorm"
)

type Guest struct {
	gorm.Model
	OrderedMeals []ordered_meal.OrderedMeal `json:"orderedMeals"`
	OrderID      uint                       `json:"-"`
}

package guest

import (
	"github.com/fdistorted/gokeeper/models/ordered-meal"
	"gorm.io/gorm"
)

type Guest struct {
	gorm.Model
	orderedMeals []ordered_meal.OrderedMeal
	OrderID      uint
}

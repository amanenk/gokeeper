package ordered_meal

import (
	"github.com/fdistorted/gokeeper/models/meal"
	"gorm.io/gorm"
)

const (
	MealOrdered = "ordered"
	MealReady   = "ready"
)

type OrderedMeal struct {
	gorm.Model
	Status  string
	MealID  uint
	Meal    meal.Meal
	GuestID uint
	BillID  uint
}

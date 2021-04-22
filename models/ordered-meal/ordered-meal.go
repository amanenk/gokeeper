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
	Status  string `json:"status"`
	MealID  uint
	Meal    meal.Meal `json:"meal"`
	GuestID uint
	BillID  uint
	OrderId uint
}

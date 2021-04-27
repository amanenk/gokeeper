package ordered_meal

import (
	"github.com/fdistorted/gokeeper/models/meal"
	"gorm.io/gorm"
)

const (
	MealOrdered = "ordered"
	MealCooking = "cooking"
	MealReady   = "ready"
)

type OrderedMeal struct {
	gorm.Model
	Status  string    `json:"status"`
	MealID  *uint     `json:"mealId"`
	Meal    meal.Meal `json:"meal" validate:"required"`
	Amount  uint      `json:"amount"`
	GuestID *uint     `json:"guestId"`
	BillID  *uint     `json:"billID"`
	OrderId uint      `json:"orderId"`
}

package bill

import (
	"github.com/fdistorted/gokeeper/models/ordered-meal"
	"gorm.io/gorm"
	"time"
)

type Bill struct {
	gorm.Model
	OrderedMeals []ordered_meal.OrderedMeal `json:"orderedMeals"`
	Total        uint                       `json:"total"`
	OrderID      uint
	CreatedAt    time.Time
}

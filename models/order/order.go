package order

import (
	"database/sql"
	"github.com/fdistorted/gokeeper/models/bill"
	"github.com/fdistorted/gokeeper/models/guest"
	orderedmeal "github.com/fdistorted/gokeeper/models/ordered-meal"
	"gorm.io/gorm"
)

const (
	StatusCreated  = "CREATED"
	StatusCanceled = "CANCELED"
	StatusBilled   = "BILLED"
	StatusFinished = "FINISHED"
)

type Order struct {
	gorm.Model
	FinishedAt   sql.NullTime              `json:"finishedAt"`
	Status       string                    `json:"status"`
	WaiterID     uint                      `validate:"required" json:"waiterId"`
	TableID      uint                      `validate:"required" json:"tableId"`
	OrderedMeals []orderedmeal.OrderedMeal `json:"orderedMeals"`
	Guests       []guest.Guest             `json:"guests"`
	Bills        []bill.Bill               `json:"bills"`
}

package order

import (
	"github.com/fdistorted/gokeeper/models/guest"
	"gorm.io/gorm"
	"time"
)

const (
	StatusStarted  = "STARTED"
	StatusCanceled = "CANCELED"
	StatusFinished = "FINISHED"
)

type Order struct {
	gorm.Model
	FinishedAt time.Time
	Status     string `validate:"required" json:"status"`
	WaiterID   uint   `validate:"required" json:"waiterId"`
	TableID    uint   `validate:"required" json:"tableId"`
	Guests     []guest.Guest
}

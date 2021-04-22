package order

import (
	"github.com/fdistorted/gokeeper/models/guest"
	"gorm.io/gorm"
	"time"
)

const (
	OrderStatusStarted  = "STARTED"
	OrderStatusCanceled = "Canceled"
	OrderStatusFinished = "Finished"
)

type Order struct {
	gorm.Model
	FinishedAt time.Time
	WaiterID   uint
	TableID    uint
	Guests     []guest.Guest
}

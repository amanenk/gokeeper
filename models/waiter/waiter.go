package waiter

import (
	"github.com/fdistorted/gokeeper/models/order"
	"gorm.io/gorm"
)

type Waiter struct {
	gorm.Model
	FirstName string
	LastName  string
	Orders    []order.Order
}

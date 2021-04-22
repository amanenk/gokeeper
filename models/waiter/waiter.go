package waiter

import (
	"github.com/fdistorted/gokeeper/models/order"
	"gorm.io/gorm"
)

type Waiter struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
	Password  string //raw password is just for example it should be hashed
	Orders    []order.Order
}

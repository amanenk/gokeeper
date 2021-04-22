package db

import (
	"github.com/fdistorted/gokeeper/config"
	"github.com/fdistorted/gokeeper/models/bill"
	"github.com/fdistorted/gokeeper/models/guest"
	"github.com/fdistorted/gokeeper/models/meal"
	"github.com/fdistorted/gokeeper/models/order"
	"github.com/fdistorted/gokeeper/models/ordered-meal"
	"github.com/fdistorted/gokeeper/models/table"
	"github.com/fdistorted/gokeeper/models/waiter"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"sync"
)

var (
	db   *gorm.DB
	once = &sync.Once{}
)

func Get() *gorm.DB {
	return db
}

func Load(cfg *config.Config) (err error) {
	once.Do(func() {
		db, err = gorm.Open(sqlite.Open(cfg.DatabaseName), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: cfg.DisableForeignKeyConstraintWhenMigrating, //for sqlite
		})
		if err != nil {
			return
		}

		// Migrate the schema
		err = db.AutoMigrate(&table.Table{},
			&bill.Bill{},
			&ordered_meal.OrderedMeal{},
			&meal.Meal{},
			&order.Order{},
			&guest.Guest{},
			&waiter.Waiter{},
		)
		if err != nil {
			return
		}
	})
	return
}

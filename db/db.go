package db

import (
	"github.com/fdistorted/gokeeper/config"
	"github.com/fdistorted/gokeeper/logger"
	"github.com/fdistorted/gokeeper/models/bill"
	"github.com/fdistorted/gokeeper/models/guest"
	"github.com/fdistorted/gokeeper/models/meal"
	"github.com/fdistorted/gokeeper/models/order"
	"github.com/fdistorted/gokeeper/models/ordered-meal"
	"github.com/fdistorted/gokeeper/models/table"
	"github.com/fdistorted/gokeeper/models/waiter"
	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(cfg *config.Config) {
	db, err := gorm.Open(sqlite.Open(cfg.DatabaseName), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: cfg.DisableForeignKeyConstraintWhenMigrating, //for sqlite
	})
	if err != nil {
		logger.Get().Fatal("failed to connect database", zap.Error(err))
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
		logger.Get().Fatal("failed to migrate the database", zap.Error(err))
	}

	DB = db
}

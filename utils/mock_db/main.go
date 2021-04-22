package main

import (
	"github.com/fdistorted/gokeeper/config"
	database "github.com/fdistorted/gokeeper/db"
	"github.com/fdistorted/gokeeper/logger"
	"github.com/fdistorted/gokeeper/models/table"
	"log"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config %+v\n", err)
	}

	err = logger.Load()
	if err != nil {
		log.Fatalf("failed to load logger %+v\n", err)
	}

	database.InitDB(cfg) // todo use db object in handlers

	// Create
	database.DB.Create(&table.Table{Number: 1, Seats: 4, IsBusy: false})
	database.DB.Create(&table.Table{Number: 2, Seats: 4, IsBusy: false})
	database.DB.Create(&table.Table{Number: 3, Seats: 2, IsBusy: false})
	database.DB.Create(&table.Table{Number: 4, Seats: 2, IsBusy: false})
	database.DB.Create(&table.Table{Number: 5, Seats: 2, IsBusy: false})
	database.DB.Create(&table.Table{Number: 6, Seats: 2, IsBusy: false})
	database.DB.Create(&table.Table{Number: 7, Seats: 2, IsBusy: false})
	database.DB.Create(&table.Table{Number: 8, Seats: 2, IsBusy: false})
	database.DB.Create(&table.Table{Number: 9, Seats: 2, IsBusy: false})
	database.DB.Create(&table.Table{Number: 10, Seats: 8, IsBusy: false})
}

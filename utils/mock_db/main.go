package main

import (
	"github.com/fdistorted/gokeeper/config"
	database "github.com/fdistorted/gokeeper/db"
	"github.com/fdistorted/gokeeper/models/meal"
	"github.com/fdistorted/gokeeper/models/table"
	"github.com/fdistorted/gokeeper/models/waiter"
	"log"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config %+v\n", err)
	}

	database.Load(cfg)

	database.Get().Create(&waiter.Waiter{FirstName: "Anna", LastName: "Smith", Email: "anna.smith123@corp1.com", Password: "anna123"})
	database.Get().Create(&waiter.Waiter{FirstName: "Jim", LastName: "Coracci", Email: "jim.c@corp1.com", Password: "jimc21111983"})
	database.Get().Create(&waiter.Waiter{FirstName: "Dennis", LastName: "Pappalois", Email: "d.pappalois@corp1.com", Password: "qwerty1@3"})

	database.Get().Create(&meal.Meal{Name: "space cake", Price: 799})
	database.Get().Create(&meal.Meal{Name: "space coffee", Price: 500})
	database.Get().Create(&meal.Meal{Name: "space burger", Price: 1299})

	database.Get().Create(&table.Table{Number: 1, Seats: 4, IsBusy: false})
	database.Get().Create(&table.Table{Number: 2, Seats: 4, IsBusy: false})
	database.Get().Create(&table.Table{Number: 3, Seats: 2, IsBusy: false})
	database.Get().Create(&table.Table{Number: 4, Seats: 2, IsBusy: false})
	database.Get().Create(&table.Table{Number: 5, Seats: 2, IsBusy: false})
	database.Get().Create(&table.Table{Number: 6, Seats: 2, IsBusy: false})
	database.Get().Create(&table.Table{Number: 7, Seats: 2, IsBusy: false})
	database.Get().Create(&table.Table{Number: 8, Seats: 2, IsBusy: false})
	database.Get().Create(&table.Table{Number: 9, Seats: 2, IsBusy: false})
	database.Get().Create(&table.Table{Number: 10, Seats: 8, IsBusy: false})
}

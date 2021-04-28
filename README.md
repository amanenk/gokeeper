# gokeeper

simple api that helps to manage a caffe

### How to run it

before testing you might want to fill the database with sample data:

```
go run ./utils/mock_db/main.go
```

Then you can run the app

```
go run main.go
```

### How to test it

Here is a collection to postman requests
https://www.getpostman.com/collections/8248743553b09c960a2d

### Task

You are creating an app to be used by servers (waiters/waitresses) at a restaurant. The app will have 3 main purposes:

* Assign guests to a particular table / server
* Take orders for the guests
* Calculate a check for the meal Create the models/schemas/APIs necessary to power this app.  
  Assumptions:  
  A system for the kitchen lets the server know when orders are ready. You do not need to address that system in your
  model. Your API should support the following scenarios:
* A group of 4 guests enters the restaurant and is assigned to a table
* Each guest makes an order for themself and a dish is ordered to share at the table
* The group splits the order into 2 bills  
  Please, send us a link to your code, with instructions to run and any additional documents you find relevant.

Basic flow:

- a few guests come to the place
- waiter checks the tables `GET http://localhost:3000/tables/?isBusy=false`
- waiter marks the table as busy `PUT http://localhost:3000/tables/2`
- waiter creates order assigned to a table`POST http://localhost:3000/orders/`
- waiter adds guests to the order `POST http://localhost:3000/orders/2/guests/`
- waiter gets the list of available meals `GET http://localhost:3000/meals`
- waiter adds ordered meal to order `POST http://localhost:3000/orders/2/order-items/`
- waiter assigns an ordered meal to the guest `POST http://localhost:3000/orders/2/guests/6/order-items/5`
- a guest leaves, waiter removes him from the order `DELETE http://localhost:3000/orders/2/guests/2/order-items/2`
- when it is time to finish the order waiter creates the bill `POST http://localhost:3000/orders/2/bills/`
- and adds ordered items to a bill `POST http://localhost:3000/orders/2/bills/1/order-items/6`
- waitier generates final bills for the order  `POST http://localhost:3000/orders/2/finish`

### TODO:

- use cache for waiters structure and for orders structure
- add websocket to notify waiters
- propagate context usage to all database transactions


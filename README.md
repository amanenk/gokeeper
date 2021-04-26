# gokeeper

simple api that helps to manage a caffe

### How to run it

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

### TODO:

- add meals endpoint
- add orderItems endpoints
- add guest endpoints
- orders endpoints
- add bills endpoint
- add websocket to notify waiters 




build:
	go build . -o server

mock_db:
	go run ./utils/mock_db/main.go
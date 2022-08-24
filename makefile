.PHONY: start

start:
	 go run "d:\Js\learn\learn-rest\cmd\apiserver\main.go"


test:
	go test -v -race -timeout 10s ./...

migrate:
	migrate -path ./migrations -database postgres://admin:admin@localhost:5432/postgres?sslmode=disable up

.DEFAULT_GOAL := start
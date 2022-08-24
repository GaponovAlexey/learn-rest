.PHONY: build

.PHONY: test
test:
	go test -v -race -timeout 10s ./...

migrate:
	migrate -path ./migrations -database postgres://admin:admin@localhost:5432/postgres?sslmode=disable up
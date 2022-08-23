.PHONY: build

.PHONY: test
test:
	go test -v -race -timeout 10s ./...

migrate:
	migrate -path ./migrations -database postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable up
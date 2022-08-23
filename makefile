.PHONY: build

.PHONY: test
test:
	go test -v -race -timeout 10s ./...
.PHONY: run
run:
	go run -v -race cmd/main.go

.PHONY: lint
## lint: check everything's okay
lint:
	golangci-lint run ./...
	go mod verify
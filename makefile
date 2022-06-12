.PHONY: run
run:
	go run -v -race main.go

.PHONY: lint
## lint: check everything's okay
lint:
	golangci-lint run ./...
	go mod verify
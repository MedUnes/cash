test:
	go run gotest.tools/gotestsum@latest --format=testdox
run:
	go run main.go
format:
	golangci-lint fmt ./...
lint:
	golangci-lint run ./...
lint-fix:
	golangci-lint run --fix ./...

.PHONY: run test format lint lint-fix
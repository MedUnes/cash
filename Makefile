test:
	go run gotest.tools/gotestsum@latest --format=testdox
benchmark:
	go test -bench=. -benchmem ./cache
run:
	go run main.go
format:
	golangci-lint fmt ./...
lint:
	golangci-lint run ./...
lint-fix:
	golangci-lint run --fix ./...
release:
	 goreleaser release --snapshot --clean --skip=publish
.PHONY: run test format lint lint-fix
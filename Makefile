test:
	go run gotest.tools/gotestsum@latest --format=testdox  -- -covermode=atomic -coverprofile=coverage.txt ./...
	grep -v "main.go" coverage.txt > coverage.tmp && mv coverage.tmp coverage.txt
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
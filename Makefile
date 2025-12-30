test:
	go run gotest.tools/gotestsum@latest --format=testdox
run:
	go run main.go
format:
	gofumpt -w .
.PHONY: run test format
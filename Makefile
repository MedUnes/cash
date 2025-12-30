test:
	go run gotest.tools/gotestsum@latest --format=testdox
run:
	go run main.go
.PHONY: run test
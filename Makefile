.PHONY: test
test:
	go test -race ./... -v

.PHONY: run 
run:
	go run main.go


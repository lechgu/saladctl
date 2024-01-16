.PHONY: clean lint

build: clean
	go build -o saladctl cmd/main.go

clean:
	rm -f saladctl

lint:
	golangci-lint run ./...
.PHONY: clean lint

build: clean
	go build -o saladctl cmd/main.go

clean:
	rm -f saladctl
	rm -rf dist

lint:
	golangci-lint run ./...

publish: clean
	mkdir -p dist/darwin
	GOOS=darwin go build -o dist/darwin/saladctl cmd/main.go
	mkdir -p dist/linux
	GOOS=linux go build -o dist/linux/saladctl cmd/main.go
	mkdir -p dist/windows
	GOOS=windows go build -o dist/windows/saladctl.exe cmd/main.go

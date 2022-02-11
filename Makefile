all: clean test build

.PHONY: build
build:
	go build -o ./bin/sudoku .

.PHONY: test
test:
	go test ./... -v -count=1

.PHONY: clean
clean:
	rm -rf ./bin

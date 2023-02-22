build:
	@go build -o bin/elevator-design

run: build
	@./bin/elevator-design

test:
	@go test -v ./...

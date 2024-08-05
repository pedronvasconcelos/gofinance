build: 
    @go build -o bin/gofinance
run: build
	@./bin/gofinance
test:
	@go test -v ./...

	

test:
	go test 

all:
	go test ./...

race:
	go test -race

cover:
	go test ./... -race -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	go tool cover -func=coverage.out

fumpt:
	@echo "Fumpting go"
	@go tool gofumpt -l -w .

lint: fumpt
	@echo "Linting go"
	@go tool golangci-lint run

bench:
	@go test -bench=.

vet:
	@go vet ./...


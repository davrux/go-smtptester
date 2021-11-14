
test:
	go test 

all:
	go test ./...

race:
	go test -race

cover:
	go test -coverprofile=coverage.out
	go tool cover -html=coverage.out

bench:
	go test -bench=.

vet:
	go vet ./...


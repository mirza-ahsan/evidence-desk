BINARY := evidence-desk
CMD    := ./cmd/evidence-desk

.PHONY: run build test vet fmt

run:
	go run $(CMD)

build:
	go build -o bin/$(BINARY) $(CMD)

test:
	go test ./...

vet:
	go vet ./...

fmt:
	gofmt -w .

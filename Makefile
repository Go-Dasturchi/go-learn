BINARY := go-learn

.PHONY: build install run fmt vet test check clean

build:
	go build -o $(BINARY) ./cmd/go-learn

install:
	go install ./cmd/go-learn

run: build
	./$(BINARY)

fmt:
	go fmt ./...

vet:
	go vet ./...

test:
	go test ./...

check: fmt vet test

clean:
	rm -f $(BINARY)

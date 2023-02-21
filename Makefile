GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test

.phony: all lib test clean update

all: lib

lib: miele/*.go
	$(GOBUILD) -v ./...

test: miele/*.go
	$(GOTEST) -v ./...

update:
	go get -u ./...
	go mod tidy

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

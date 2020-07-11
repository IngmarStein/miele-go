GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test

.phony: all lib test clean

all: lib

lib: miele/*.go
	$(GOBUILD) -v ./...

test: miele/*.go
	$(GOTEST) -v ./...

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

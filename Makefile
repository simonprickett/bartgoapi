GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean

DEPS_FOLDER=vendor

BINARY_FOLDER=dist
BINARY_NAME=bart-json-api
BINARY_REL_NAME=$(BINARY_FOLDER)/$(BINARY_NAME)

all: build

build: ensure
	$(GOBUILD) -o $(BINARY_REL_NAME)

run: build
	$(BINARY_REL_NAME)

ensure:
	dep ensure

clean:
	rm -rf $(BINARY_FOLDER)
	rm -rf $(DEPS_FOLDER)

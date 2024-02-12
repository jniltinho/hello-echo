GOCMD ?= go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOGET = $(GOCMD) get
BINARY_NAME = hello-echo


default: get generate build


get:
	@go mod tidy

generate:
	@templ generate


build:
	CGO_ENABLED=0 $(GOBUILD) -o $(BINARY_NAME) -v -ldflags="-s -w"


clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)


run:
	./$(BINARY_NAME)
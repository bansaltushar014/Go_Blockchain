GOCMD := go
GOBUILD := $(GOCMD) build
GOCLEAN := $(GOCMD) clean
GOTEST := $(GOCMD) test
GOGET := $(GOCMD) get

build:
	$(GOBUILD) -o main main.go

test:
	$(GOTEST) ./...

clean:
	$(GOCLEAN)
	rm -f main

run:
	$(GOBUILD) -o main main.go
	./main
GOTEST=go test
GOGET=go get


all: test

test:
	$(GOTEST) -v ./problems

deps:
	$(GOGET) github.com/stretchr/testify/assert


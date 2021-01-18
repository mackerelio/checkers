.PHONY: all
all: clean test

.PHONY: test
test:
	go test $(TESTFLAGS) ./...

.PHONY: cover
cover:
	go test -race -covermode atomic -coverprofile=profile.cov ./...

.PHONY: clean
clean:
	go clean

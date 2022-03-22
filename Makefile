.PHONY: all
all: clean test

.PHONY: test
test:
	go test $(TESTFLAGS) ./...

.PHONY: clean
clean:
	go clean

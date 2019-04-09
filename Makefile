.PHONY: all
all: clean test

.PHONY: test
test: lint
	go test $(TESTFLAGS) ./...

.PHONY: deps
deps:
	go get -d -v -t ./...

.PHONY: devel-deps
devel-deps: deps
	GO111MODULE=off \
	go get golang.org/x/lint/golint \
		github.com/axw/gocov/gocov \
		github.com/mattn/goveralls

.PHONY: lint
lint: devel-deps
	go vet ./...
	golint -set_exit_status ./...

.PHONY: cover
cover: devel-deps
	tool/cover.sh

.PHONY: clean
clean:
	go clean

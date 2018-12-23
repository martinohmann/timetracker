.DEFAULT_GOAL := help

VERSION ?= "git-$(shell git rev-parse --short=7 HEAD)"
BUILD_TIME := $(shell date)

.PHONY: help
help:
	@grep -E '^[a-zA-Z-]+:.*?## .*$$' Makefile | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "[32m%-12s[0m %s\n", $$1, $$2}'

.PHONY: build
build: ## build
	go build -o timetracker \
		-ldflags="-X 'github.com/martinohmann/timetracker/pkg/version.Version=$(VERSION)' \
			-X 'github.com/martinohmann/timetracker/pkg/version.BuildTime=$(BUILD_TIME)'" \
		main.go

.PHONY: test
test: ## run tests
	go test -race -tags="$(TAGS)" $$(go list ./... | grep -v /vendor/)

.PHONY: vet
vet: ## run go vet
	go vet $$(go list ./... | grep -v /vendor/)

.PHONY: coverage
coverage: ## generate code coverage
	scripts/coverage

.PHONY: misspell
misspell: ## check spelling in go files
	misspell *.go


.PHONY: clean
clean: ## clean artifacts and dependencies
	rm -rf vendor/ ttc

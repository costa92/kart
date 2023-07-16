.DEFAULT_GOAL := all

.PHONY: all
all: tidy gen add-copyright format lint cover build


## lint: Check syntax and styling of go sources.
.PHONY: lint
lint:
	golangci-lint version
	golangci-lint run -v --color always --out-format colored-line-number
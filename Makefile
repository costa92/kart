.DEFAULT_GOAL := all

.PHONY: all
all: tidy gen add-copyright format lint cover build


## lint: Check syntax and styling of go sources.
.PHONY: lint
lint:
	@$(MAKE) go.lint
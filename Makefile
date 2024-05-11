IS_IN_PROGRESS = "is in progress ..."


## build-cli: build the go migrate cli
.PHONY: build-cli
build-cli:
	@echo "make build-cli ${IS_IN_PROGRESS}"
	@go build -o mcli
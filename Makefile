NOW = $(shell date -u '+%Y%m%d%I%M%S')

RELEASE_VERSION = v1.0.0

APP 			= rizhua
SERVER_BIN  	= ./bin/${APP}
RELEASE_ROOT 	= ./bin/release
RELEASE_SERVER 	= ${RELEASE_ROOT}/${APP}
GIT_COUNT 		= $(shell git rev-list --all --count)
GIT_HASH        = $(shell git rev-parse --short HEAD)
RELEASE_TAG     = $(RELEASE_VERSION).$(GIT_COUNT).$(GIT_HASH)

all: start

.PHONY: build
build:
	@go build -ldflags "-w -s -X main.VERSION=$(RELEASE_TAG)" -o $(SERVER_BIN) ./cmd

.PHONY: start
start:
	@go run -ldflags "-X main.VERSION=$(RELEASE_TAG)" ./cmd/main.go web -c ./infrastructure/etc/config.toml -m ./infrastructure/etc/model.conf

.PHONY: swagger
swagger:
	@hash swag > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		go install github.com/swaggo/swag/cmd/swag@latest; \
	fi
	@swag init --generalInfo ./cmd/main.go --output ./interface/swagger

.PHONY: wire
wire:
	@hash wire > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		go install github.com/google/wire/cmd/wire@latest; \
	fi
	@wire gen ./infrastructure/injector

.PHONY: clean
clean:
	rm -rf $(SERVER_BIN)
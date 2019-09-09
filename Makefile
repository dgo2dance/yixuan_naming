PROJECT=yixuan_naming
PREFIX=$(shell pwd)
VENDOR=vendor
BIN=naming

# Env
ifndef GO
	GO=/usr/bin/go
endif

ifndef GOFMT
	GOFMT=/usr/bin/gofmt
endif

.PHONY: all clean install uninstall
.DEFAULT: all

# Targets
all: fmt build

fmt:
	@echo -e "\033[1;32m * Source code format checking ...\033[0m"
	@echo -e "\033[1;37m   @ gofmt\033[0m source code"
	@test -z "$$(find . -name \"*.go\" -not -path \"$(VENDOR)/*\" -exec $(GOFMT) -s -l '{}' +)"

build:
	@echo -e "\033[1;33m + Building ${PROJECT} ...\033[0m"
	@mkdir -p ./bin
	@echo -e "\033[1;37m   @ `$(GO) version`\033[0m"
	@echo -e "\033[0;34m     + $(BIN)\033[0m"
	@GO11MODULE=on CGO_ENABLED=0 GOOS=linux $(GO) build -a -ldflags '-extldflags "-static"' -o ./bin/$(BIN)

clean:
	@echo
	@echo -e "\033[1;35m - Cleaning ${PROJECT} ...\033[0m"
	@rm ./bin/* -f
	@rm $(VENDOR)/* -rf
	@echo
	@echo

install:
	@echo
	@echo "\033[1;34m + Installing ${PROJECT} ...\033[0m"
	@cp ./bin/* /usr/local/bin
	@echo
	@echo

uninstall:
	@echo
	@echo "\033[1;33m - Uninstall ${PROJECT} ...\033[0m"
	@rm /usr/local/bin/$(BIN) -f
	@echo
	@echo

dockcerfile:
	@echo


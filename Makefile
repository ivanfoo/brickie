# Default shell
SHELL := /bin/bash

# General
WORKDIR = $(PWD)

# Go parameters
GOCMD = go
GOTEST = $(GOCMD) test -v

# Coverage
COVERAGE_REPORT = coverage.txt
COVERAGE_PROFILE = profile.out
COVERAGE_MODE = atomic

# Environment
BUILD_PATH := $(TRAVIS_BUILD_DIR)/build
BUILD := $(shell date +"%m-%d-%Y_%H_%M_%S")
COMMIT := $(shell git log --format='%H' -n 1 | cut -c1-10)
BRANCH := $(shell git rev-parse --abbrev-ref HEAD)

# Packages content
PKG_OS = darwin linux
PKG_ARCH = amd64

ifneq ($(origin CI), undefined)
    WORKDIR := $(TRAVIS_BUILD_DIR)
endif


test:
    $(GOTEST) ./...

packages:
    for os in $(PKG_OS); do \
        for arch in $(PKG_ARCH); do \
            cd $(TRAVIS_BUILD_DIR); \
            mkdir -p $(BUILD_PATH)/$(PROJECT)_$${os}_$${arch}; \
            for cmd in $(COMMANDS); do \
                if [ -d "$${cmd}" ]; then \
                    cd $${cmd}; \
                fi; \
                GOOS=$${os} GOARCH=$${arch} $(GOCMD) build -ldflags \
                "-X main.version=$(BRANCH) -X main.build=$(BUILD) -X main.commit=$(COMMIT)" \
                -o "$(BUILD_PATH)/$(PROJECT)_$${os}_$${arch}/`basename $${PWD}`" .; \
                cd $(TRAVIS_BUILD_DIR); \
            done; \
            cd $(BUILD_PATH); \
            tar -cvzf $(PROJECT)_$(BRANCH)_$${os}_$${arch}.tar.gz $(PROJECT)_$${os}_$${arch}/; \
        done; \
    done; \

test-coverage:
    cd $(WORKDIR); \
    echo "" > $(COVERAGE_REPORT); \
    for dir in `find . -name "*.go" | grep -o '.*/' | sort | uniq`; do \
        $(GOTEST) $$dir -coverprofile=$(COVERAGE_PROFILE) -covermode=$(COVERAGE_MODE); \
        if [ $$? != 0 ]; then \
            exit 2; \
        fi; \
        if [ -f $(COVERAGE_PROFILE) ]; then \
            cat $(COVERAGE_PROFILE) >> $(COVERAGE_REPORT); \
            rm $(COVERAGE_PROFILE); \
        fi; \
    done; \

# Copyright (c) 2022 Intel Corporation.
# SPDX-License-Identifier: Apache-2.0

# local binary builds
WORKSPACE    = _workspace
BINDIR       = $(WORKSPACE)/bin
SERVER     	 = server 

GOCMD    := GOPRIVATE="github.com/intel-innersource/*" go


.PHONY: all go-tidy go-build clean help
all: go-build

$(BINDIR):
	mkdir -p $@

go-tidy:
	$(GOCMD) mod tidy

go-build: $(BINDIR) ## build binaries
	$(GOCMD) build -v -o $(BINDIR)/$(SERVER) cmd/server/server.go

clean: ## clean up
	rm -rf $(BINDIR)

help:
	@echo harvester-vm manager make targets
	@echo
	@grep '^[[:alnum:]_-]*:.* ##' $(MAKEFILE_LIST) \
    | sort | awk 'BEGIN {FS=":.* ## "}; {printf "%-25s %s\n", $$1, $$2};'
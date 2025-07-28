# SPDX-FileCopyrightText: 2025 James Pond <james@cipher.host>
#
# SPDX-License-Identifier: EUPL-1.2

.POSIX:
.SUFFIXES:

PKGNAME := taskemon
PKGBIN  := $(PKGNAME)

PREFIX     := /usr/local
BINDIR     := bin
PKGDIR     := ./
OUTDIR     := build

GO        ?= go
REUSE     ?= reuse
RM        ?= rm
INSTALL   ?= install

GOBUILD_FLAGS := -trimpath \
				 -buildmode=pie \
				 -mod=readonly \
				 -modcacherw
LDFLAGS       := -s -w
GOBUILD_OUT   := -o $(OUTDIR)/$(PKGBIN)
GOBUILD_OPTS  := $(GOBUILD_FLAGS) -ldflags '$(LDFLAGS)' $(GOBUILD_OUT)

all: build

build: # Builds an application binary.
	$(GO) build $(GOBUILD_OPTS) $(PKGDIR)

install: # Installs the release binary.
	$(INSTALL) -d \
		$(DESTDIR)$(PREFIX)/$(BINDIR)/
	$(INSTALL) -pm 0755 $(OUTDIR)/$(PKGBIN) $(DESTDIR)$(PREFIX)/$(BINDIR)/

lint: # Runs golangci-lint using the config at the root of the repository.
	$(GO) run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest run ./...

lint/licenses: # Run reuse to ensure the project complies with the REUSE specification.
	$(REUSE) lint --lines

licenses: # Runs go-licenses to check the licenses of the dependencies and generate a CSV file.
	$(GO) run github.com/google/go-licenses@latest report \
		--template '.github/license-3rdparty.tpl' \
		--ignore 'git.sr.ht/~jamesponddotco/taskemon' \
		'git.sr.ht/~jamesponddotco/taskemon' > LICENSE-3rdparty.csv

clean: # Cleans cache files from tests and deletes any build output.
	$(RM) -rf $(OUTDIR)

.PHONY: all build install lint lint/licenses licenses clean

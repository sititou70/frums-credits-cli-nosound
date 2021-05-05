GO = go
BINNAME = credits
BIN_DIR = dist/
REPONAME = frums-credits-cli

NATIVE_BIN=$(BIN_DIR)$(BINNAME)
native: tidy $(NATIVE_BIN)
tidy:
	$(GO) mod tidy
$(NATIVE_BIN): $(shell find . -type f -name '*.go')
	$(GO) build -o $(NATIVE_BIN) -v

build-image:
	docker build -t frums-credits-cli-build .
cross: build-image
	docker run --rm --privileged \
  	-v $(CURDIR):/$(REPONAME) \
  	-w /$(REPONAME) \
  	frums-credits-cli-build --snapshot --skip-publish --rm-dist

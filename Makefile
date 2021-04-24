GO = go
BINNAME = credits

$(BINNAME): $(shell find . -type f -name '*.go')
	$(GO) build -o $(BINNAME) -v

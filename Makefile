BINNAME=rohan
SRVNAME=gondor
RELEASE=-s -w
UPXBIN=/usr/local/bin/upx
GOBIN=/usr/local/bin/go
GOOS=$(shell uname -s | tr [A-Z] [a-z])
GOARGS=GOARCH=amd64 CGO_ENABLED=0
GOBUILD=$(GOARGS) $(GOBIN) build -ldflags="$(RELEASE)"

.PHONY: all
all: clean build
build:
	@echo "Compile $(BINNAME) ..."
	GOOS=$(GOOS) $(GOBUILD) -o $(BINNAME) ./cmd/rohan/
	@echo "Build success."
clean:
	rm -f $(BINNAME)
	@echo "Clean all."
upx: build command
	$(UPXBIN) $(BINNAME) $(SRVNAME)
upxx: build command
	$(UPXBIN) --ultra-brute $(BINNAME) $(SRVNAME)
serv:
	@echo "Compile $(SRVNAME) ..."
	GOOS=$(GOOS) $(GOBUILD) -o $(SRVNAME) ./cmd/gondor/
	@echo "Build success."
vend:
	GOOS=$(GOOS) $(GOBUILD) -mod=vendor -o $(BINNAME) ./cmd/rohan/
	GOOS=$(GOOS) $(GOBUILD) -mod=vendor -o $(SRVNAME) ./cmd/gondor/

BINNAME=rohan
CMDNAME=moria
SRVNAME=gondor
RELEASE=-s -w
UPXBIN=/usr/local/bin/upx
GOBIN=/usr/local/bin/go
GOOS=$(shell uname -s | tr [A-Z] [a-z])
GOARGS=GOARCH=amd64 CGO_ENABLED=1
GOBUILD=$(GOARGS) $(GOBIN) build -ldflags="$(RELEASE)"

.PHONY: rebuild
all: clean serv build
build:
	@echo "Compile $(BINNAME) $(CMDNAME) ..."
	GOOS=$(GOOS) $(GOBUILD) -o $(BINNAME) ./cmd/rohan/
	GOOS=$(GOOS) $(GOBUILD) -o $(CMDNAME) ./cmd/moria/
	@echo "Build success."
clean:
	rm -f $(BINNAME) $(CMDNAME)
	@echo "Clean all."
upx: build command
	$(UPXBIN) $(BINNAME) $(CMDNAME) $(SRVNAME)
upxx: build command
	$(UPXBIN) --ultra-brute $(BINNAME) $(CMDNAME) $(SRVNAME)
rebuild: clean build
serv:
	@echo "Remove and Compile $(SRVNAME) ..."
	rm -f $(SRVNAME)
	GOOS=$(GOOS) $(GOBUILD) -o $(SRVNAME) ./cmd/gondor/
	@echo "Build success."
vend:
	GOOS=$(GOOS) $(GOBUILD) -mod=vendor -o $(BINNAME) ./cmd/rohan/
	GOOS=$(GOOS) $(GOBUILD) -mod=vendor -o $(CMDNAME) ./cmd/moria/
	GOOS=$(GOOS) $(GOBUILD) -mod=vendor -o $(SRVNAME) ./cmd/gondor/

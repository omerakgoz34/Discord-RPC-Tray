RELEASE_FLAGS := -ldflags="-s -w -H=windowsgui"
DEBUG_FLAGS := -ldflags="" -tags=debug
BINARY_NAME := Discord-RPC-Tray.exe

all: build

build:
	go build -v $(DEBUG_FLAGS) -o $(BINARY_NAME)

run: build
	$(BINARY_NAME)

release:
	go build -v $(RELEASE_FLAGS) -o $(BINARY_NAME)

clean:
ifeq ($(OS),Windows_NT)
	del /Q $(BINARY_NAME)
else
	rm -r $(BINARY_NAME)
endif
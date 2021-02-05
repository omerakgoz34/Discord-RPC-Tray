PROD_FLAGS := -ldflags="-s -w"
DEBUG_FLAGS := -ldflags="" -tags=debug
BINARY_NAME := Discord-RPC-Tray.exe

all: build

build:
	go build -v $(DEBUG_FLAGS) -o $(BINARY_NAME)

run: build
	$(BINARY_NAME)

production:
	go build -v $(PROD_FLAGS) -o $(BINARY_NAME)

clean:
ifeq ($(OS),Windows_NT)
	del /Q $(BINARY_NAME)
else
	rm -r $(BINARY_NAME)
endif
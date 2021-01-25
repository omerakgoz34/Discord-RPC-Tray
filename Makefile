BUILD_FLAGS := -ldflags="-s -w"
BINARY_NAME := Discord-RPC-Tray.exe

all: build

build:
	go build -v $(BUILD_FLAGS) -o $(BINARY_NAME)

run: build
	$(BINARY_NAME)

clean:
ifeq ($(OS),Windows_NT)
	del /Q $(BINARY_NAME)
else
	rm -r $(BINARY_NAME)
endif
all: build

build: clean
	go build -v -ldflags="" -tags=debug

release: clean
ifeq ($(OS),Windows_NT)
	go build -v -ldflags="-s -w -H=windowsgui"
else
	go build -v -ldflags="-s -w"
endif

run:
	$(BINARY_NAME)

clean:
ifeq ($(OS),Windows_NT)
	del /Q Discord-RPC-Tray.exe
else
	rm -rf Discord-RPC-Tray
endif
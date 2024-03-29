all: build

build:
	go build -v -ldflags="" -tags=debug

release:
ifeq ($(OS),Windows_NT)
	go build -v -ldflags="-s -w -H=windowsgui"
else
	go build -v -ldflags="-s -w"
endif

run:
ifeq ($(OS),Windows_NT)
	.\Discord-RPC-Tray.exe
else
	./Discord-RPC-Tray
endif

clean:
ifeq ($(OS),Windows_NT)
	del /Q Discord-RPC-Tray.exe
else
	rm -rf Discord-RPC-Tray
endif
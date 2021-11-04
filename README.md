# Discord-RPC-Tray

Lang: [EN] [TR](/README_TR.md)

Most lightweight app to make custom Discord "Playing" statuses!  
  
[![BUILD: Windows](https://github.com/omerakgoz34/Discord-RPC-Tray/actions/workflows/build-windows.yml/badge.svg)](https://github.com/omerakgoz34/Discord-RPC-Tray/actions/workflows/build-windows.yml)
[![BUILD: Linux](https://github.com/omerakgoz34/Discord-RPC-Tray/actions/workflows/build-linux.yml/badge.svg)](https://github.com/omerakgoz34/Discord-RPC-Tray/actions/workflows/build-linux.yml)
[![BUILD: macOS](https://github.com/omerakgoz34/Discord-RPC-Tray/actions/workflows/build-macos.yml/badge.svg)](https://github.com/omerakgoz34/Discord-RPC-Tray/actions/workflows/build-macos.yml)  

## Runtime Requirements
* Windows: Windows 7 or newer
* Linux: GTK+ v3.10 or newer
* macOS: v10.8 or newer

## Downloads (v1.1.0)
* [Windows x64](https://github.com/omerakgoz34/Discord-RPC-Tray/releases/download/v1.1.0/Discord-RPC-Tray_v1.1.0_win64.zip) ([scoop file](https://github.com/omerakgoz34/Discord-RPC-Tray/blob/d3f35ba40ab0b3c6ea2e9b9918b2135e247ee501/discord-rpc-tray.json))
* [Linux x64](https://github.com/omerakgoz34/Discord-RPC-Tray/releases/download/v1.1.0/Discord-RPC-Tray_v1.1.0_linux64.zip) (not tested)
* [macOS x64](https://github.com/omerakgoz34/Discord-RPC-Tray/releases/download/v1.1.0/Discord-RPC-Tray_v1.1.0_macos64.app.zip) (not tested) (thanks to [@elvodqa](https://github.com/elvodqa))

## Screenshots
![Screenshot_66](https://user-images.githubusercontent.com/49201485/140165938-701e88ab-fd12-4560-ad39-a5b6cf5560c1.png)  
![Screenshot_67](https://user-images.githubusercontent.com/49201485/140166368-ade1880a-68f0-4ea7-8b46-2738f9851a2e.png)  
![Screenshot_62](https://user-images.githubusercontent.com/49201485/140166003-c275fa33-aa40-4bd5-93c4-590ade3488b1.png)

## Features
* Super low resource usage! (0% CPU and ~4MB RAM...)
* All rich-presence features can useable. (Status updated every 12 seconds.)
* You can use your own [discord application](https://discord.com/developers/applications) for custom app name and images.
* Dynamic icon for indicating RPC is active or not. (Green dot when RPC is active and red dot when not active.)
* A button for getting the current timestamp easily.

## Go timestamp format (Timestamp fields)
2021-11-03 17:40:20.6396501 +0300 +03  

* 2021-11-03 --> date
* 17:40:20 --> time (hours:minutes:seconds)
* .6396501 --> miliseconds (who uses dis for time except computers :p)
* +0300 --> Time Zone. Use Z for +00:00 (probably it's already automatically set to your system UTC setting.)

## Used libraries
* Tray icon: [getlantern/systray](https://github.com/getlantern/systray)
* GUI (graphical user interface): [andlabs/ui](https://github.com/github.com/andlabs/ui)
* Discord RPC (rich-presence): [hugolgst/rich-go](https://github.com/hugolgst/rich-go)
* Opening files and URLs with default programs: [skratchdot/open-golang](https://github.com/skratchdot/open-golang)
* Copying timestamp to clipboard: [atotto/clipboard](https://github.com/atotto/clipboard)

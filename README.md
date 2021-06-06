# Discord-RPC-Tray

Most lightweight app to make custom Discord "Playing" statuses :P  
  
[![BUILD: Windows](https://github.com/omerakgoz34/Discord-RPC-Tray/actions/workflows/build-windows.yml/badge.svg)](https://github.com/omerakgoz34/Discord-RPC-Tray/actions/workflows/build-windows.yml)
[![BUILD: Linux](https://github.com/omerakgoz34/Discord-RPC-Tray/actions/workflows/build-linux.yml/badge.svg)](https://github.com/omerakgoz34/Discord-RPC-Tray/actions/workflows/build-linux.yml)
[![BUILD: macOS](https://github.com/omerakgoz34/Discord-RPC-Tray/actions/workflows/build-macos.yml/badge.svg)](https://github.com/omerakgoz34/Discord-RPC-Tray/actions/workflows/build-macos.yml)  

## Screenshots  
![Screenshot_240](https://user-images.githubusercontent.com/49201485/120932531-e7ed1800-c6fe-11eb-9d3b-dd016403f6df.png)  
![Screenshot_237](https://user-images.githubusercontent.com/49201485/120929660-8757de00-c6f2-11eb-87b8-74cbab6ecb02.png)  
![Screenshot_238](https://user-images.githubusercontent.com/49201485/120929803-2b418980-c6f3-11eb-8fd2-7598656fe9ec.png)  

## Features
* Super low resource usage!
* No UI. All settings can be changed with an easily accessible config file.
* All rich-presence features can useable. (Status updated every 12 seconds.)
* You can use your own [discord application](https://discord.com/developers/applications) for custom app name and images.
* Dynamic icon for indicating RPC is active or not. (Green dot when RPC is active and red dot when not active.)
* A button for getting the current timestamp easily.

## Usage
* After changing config, need to reload config.
* Also there is a sample config file accessible from tray menu for referencing.
* If any error occurs, the app closes itself automatically. Because there is no UI to notify the user :p
* For using "elapsed"(Start) and "left"(End) time indicators (you can only use one of them at the same time), you can use the value of the DateNow(read-only) field in the config: `"DateNow": "2021-06-06T19:32:50.3947031+03:00"`

## Go timestamp format (DateNow and Timestamp.Start/End fields)
2021-06-06 T 19:32:50 .3947031 +03:00  
2021-06-06 T 16:32:50 .3947031 Z
* 2021-06-06 --> date
* 19:32:50 --> time
* .3947031 --> miliseconds (no need to change)
* +03:00 --> UTC. Use Z for +00:00 (probably it's already automatically set to your system UTC setting.)

## Used libraries
* Tray icon: [github.com/getlantern/systray](https://github.com/getlantern/systray)
* Discord RPC(rich-presence): [github.com/hugolgst/rich-go](https://github.com/hugolgst/rich-go)
* Opening files and URLs with default programs: [github.com/skratchdot/open-golang](https://github.com/skratchdot/open-golang)

# Discord-RPC-Tray

Most lightweight app to make custom Discord "playing" statuses :P (WIP)  
[![BUILD: Windows](https://github.com/omerakgoz34/Discord-RPC-Tray/actions/workflows/build_windows.yml/badge.svg)](https://github.com/omerakgoz34/Discord-RPC-Tray/actions/workflows/build_windows.yml)
[![BUILD: Linux](https://github.com/omerakgoz34/Discord-RPC-Tray/actions/workflows/build_linux.yml/badge.svg)](https://github.com/omerakgoz34/Discord-RPC-Tray/actions/workflows/build_linux.yml)

## Installation
* **Windows:** Download the exe file from the [Releases](https://github.com/omerakgoz34/Discord-RPC-Tray/releases) page and run it. (I will add it to [Scoop](https://scoop.sh) extras repo in the future)
* **Linux:** *WIP...*
* **macOS:** *WIP...*

## Features
* Written in Go :P
* Terminal based UI for fast and lightweight usage
* Uses 0,0% CPU and 4-8MB of RAM on my windows PC
* Of course a Tray icon for hiding the console (for now it's uses default icon until I make a new one :p)
* UPX compression for smaller file sizes

Maybe in the future I can add some extra features like displaying the CPU and RAM usage of the PC on Discord profile :D

## Screenshots
(For now only Details and State fields usable)  
![Screenshot_146](https://user-images.githubusercontent.com/49201485/114171439-f60bfb80-993c-11eb-8ee8-44ac8ea3e870.png)  
![Screenshot_147](https://user-images.githubusercontent.com/49201485/114171507-0ae88f00-993d-11eb-8719-6e92439e1c70.png)

## Used Libraries
* [rivo/tview](https://github.com/rivo/tview) for terminal ui
* [getlantern/systray](https://github.com/getlantern/systray) for tray icon control
* [hugolgst/rich-go](https://github.com/hugolgst/rich-go) for Discord RPC connection

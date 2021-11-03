# Discord-RPC-Tray

Dil: [EN](/README.md) [TR]

Özel Discord "Oynuyor" durumu yapmak için en hafif uygulama!  
  
[![BUILD: Windows](https://github.com/omerakgoz34/Discord-RPC-Tray/actions/workflows/build-windows.yml/badge.svg)](https://github.com/omerakgoz34/Discord-RPC-Tray/actions/workflows/build-windows.yml)
[![BUILD: Linux](https://github.com/omerakgoz34/Discord-RPC-Tray/actions/workflows/build-linux.yml/badge.svg)](https://github.com/omerakgoz34/Discord-RPC-Tray/actions/workflows/build-linux.yml)
[![BUILD: macOS](https://github.com/omerakgoz34/Discord-RPC-Tray/actions/workflows/build-macos.yml/badge.svg)](https://github.com/omerakgoz34/Discord-RPC-Tray/actions/workflows/build-macos.yml)  

## İndirmeler (v1.1.0)

* [Windows x64](https://github.com/omerakgoz34/Discord-RPC-Tray/releases/download/v1.1.0/Discord-RPC-Tray_v1.1.0_win64.zip) ([scoop dosyası](https://github.com/omerakgoz34/Discord-RPC-Tray/blob/cc2e6e6a68c8906311bb091a61842fa71811373e/discord-rpc-tray.json))
* [Linux x64](https://github.com/omerakgoz34/Discord-RPC-Tray/releases/download/v1.1.0/Discord-RPC-Tray_v1.1.0_linux64.zip) (test edilmedi)
* [macOS x64](https://github.com/omerakgoz34/Discord-RPC-Tray/releases/download/v1.1.0/Discord-RPC-Tray_v1.1.0_macos64.app.zip) (test edilmedi) ([@elvodqa](https://github.com/elvodqa)'ya teşekkürler)

## Ekran Görüntüleri

![Screenshot_66](https://user-images.githubusercontent.com/49201485/140165938-701e88ab-fd12-4560-ad39-a5b6cf5560c1.png)  
![Screenshot_67](https://user-images.githubusercontent.com/49201485/140166368-ade1880a-68f0-4ea7-8b46-2738f9851a2e.png)  
![Screenshot_62](https://user-images.githubusercontent.com/49201485/140166003-c275fa33-aa40-4bd5-93c4-590ade3488b1.png)

## Özellikler

* Süper düşük kaynak kullanımı! (0% CPU ve ~4MB RAM...)
* Bütün rich-presence özellikleri kullanılabilir. (Durum her 12 saniyede bir güncellenir.)
* Kendi [discord uygulamanızı](https://discord.com/developers/applications) kullanarak özel uygulama adı ve resmi yapabilirsiniz.
* RPC'nin aktif olup olmadığını gösteren dinamik simge. (RPC aktif iken yeşil nokta ve aktif değilken kırmızı nokta)
* Kolayca timestamp almayı sağlayan bir buton.

## Go timestamp formatı (Timestamp bölümleri)

2021-11-03 17:40:20.6396501 +0300 +03  

* 2021-11-03 --> tarih
* 17:40:20 --> zaman (saat:dakika:saniye)
* .6396501 --> milisaniye (bunu bilgisayarlar dışında kim saat için kullanır ki :p)
* +0300 --> UTC. +00:00 için Z kullanınız. (muhtemelen otomatik olarak sisteminizdeki UTC zaman ayarıyla aynı ayarlanmıştır.)

## Kullanılan kütüphaneler

* Tray icon: [getlantern/systray](https://github.com/getlantern/systray)
* GUI (grafiksel kullanıcı arayüzü): [andlabs/ui](https://github.com/github.com/andlabs/ui)
* Discord RPC (rich-presence): [hugolgst/rich-go](https://github.com/hugolgst/rich-go)
* Dosyaları ve URL'leri varsayılan uygulamalarla açma: [skratchdot/open-golang](https://github.com/skratchdot/open-golang)
* Timestamp'i panoya kopyalamak için: [atotto/clipboard](https://github.com/atotto/clipboard)

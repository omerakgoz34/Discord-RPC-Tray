# Discord-RPC-Tray

[İNGİLİZCE](/README.md) | TÜRKÇE

Özel Discord "Oynuyor" durumu yapmak için en hafif uygulama!  
  
[![BUILD: Windows](https://github.com/omerakgoz34/Discord-RPC-Tray/actions/workflows/build-windows.yml/badge.svg)](https://github.com/omerakgoz34/Discord-RPC-Tray/actions/workflows/build-windows.yml)
[![BUILD: Linux](https://github.com/omerakgoz34/Discord-RPC-Tray/actions/workflows/build-linux.yml/badge.svg)](https://github.com/omerakgoz34/Discord-RPC-Tray/actions/workflows/build-linux.yml)
[![BUILD: macOS](https://github.com/omerakgoz34/Discord-RPC-Tray/actions/workflows/build-macos.yml/badge.svg)](https://github.com/omerakgoz34/Discord-RPC-Tray/actions/workflows/build-macos.yml)  

## İndirme Linkleri (v1.0.0)

* [Windows x64](https://github.com/omerakgoz34/Discord-RPC-Tray/releases/download/v1.0.0/Discord-RPC-Tray_v1.0.0_windows64.zip)
* Linux x64 (test ediliyor)
* [macOS x64](https://github.com/omerakgoz34/Discord-RPC-Tray/releases/download/v1.0.0/Discord-RPC-Tray_v1.0.0_macos64.app.zip) ([@elvodqa](https://github.com/elvodqa))

## Ekran Görüntüleri

![Screenshot_240](https://user-images.githubusercontent.com/49201485/120932531-e7ed1800-c6fe-11eb-9d3b-dd016403f6df.png)  
![Screenshot_237](https://user-images.githubusercontent.com/49201485/120929660-8757de00-c6f2-11eb-87b8-74cbab6ecb02.png)  
![Screenshot_238](https://user-images.githubusercontent.com/49201485/120929803-2b418980-c6f3-11eb-8fd2-7598656fe9ec.png)  

## Özellikler

* Süper düşük kaynak kullanımı! (0% CPU ve ~3MB RAM...)
* Arayüz yok. Bütün ayarlar kolayca ulaşılabilen bir config dosyasıyla değiştirilebilir.
* Bütün rich-presence özellikleri kullanılabilir. (Durum her 12 saniyede bir güncellenir.)
* Kendi [discord uygulamanızı](https://discord.com/developers/applications) kullanarak özel uygulama adı ve resmi yapabilirsiniz.
* RPC'nin aktif olup olmadığını gösteren dinamik simge. (RPC aktif iken yeşil nokta ve aktif değilken kırmızı nokta)
* Kolayca o anki timestamp'i almayı sağlayan bir buton.

## Kullanım

* Config yolu windows için: `C:\Users\omerakgoz34\AppData\Roaming\Discord RPC Tray\config.json` linux ve macos için: `~/.config/Discord RPC Tray/config.json`
* Config değiştirildikten sonra reload yapılması gerek.
* Ayrıca referans almak için tray menüsünde örnek bir config dosyası mevcut.
* Herhangi bir hata meydana gelirse uygulama otomatik olarak kendini kapatır. Çünkü hatayı kullanıcıya bildirecek bir arayüz yok :p
* "geçti"(Start) ve "kaldı"(End) zaman göstergelerini kullanabilmek için (aynı anda sadece birini kullanabilirsiniz), config dosyasındaki DateNow(sadece-okunabilir) değerini kullanabilirsiniz: `"DateNow": "2021-06-06T19:32:50.3947031+03:00"`

## Go timestamp formatı (DateNow ve Timestamp.Start/End bölümleri)

2021-06-06 T 19:32:50 .3947031 +03:00  
2021-06-06 T 16:32:50 .3947031 Z

* 2021-06-06 --> tarih
* 19:32:50 --> zaman
* .3947031 --> milisaniye (değiştirmeye gerek yok)
* +03:00 --> UTC. +00:00 için Z kullanınız. (muhtemelen otomatik olarak sisteminizdeki UTC zaman ayarıyla aynı ayarlanmıştır.)

## Kullanılan kütüphaneler

* Tray icon: [getlantern/systray](https://github.com/getlantern/systray)
* Discord RPC (rich-presence): [hugolgst/rich-go](https://github.com/hugolgst/rich-go)
* Dosyaları ve URL'leri varsayılan uygulamalarla açma: [skratchdot/open-golang](https://github.com/skratchdot/open-golang)

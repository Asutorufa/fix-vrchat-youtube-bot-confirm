#

## yt-dlp wiki

<https://github.com/yt-dlp/yt-dlp/wiki/FAQ#how-do-i-pass-cookies-to-yt-dlp>

## how to use

### 1. go to vr chat yt-dlp directory

open powershell or linux terminal

windows at: `[YOU USER DIR]/AppData/LocalLow/VRChat/VRChat/Tools`
linux at: `[STEAM LIBRARY DIR]/steamapps/compatdata/438100/pfx/drive_c/users/steamuser/AppData/LocalLow/VRChat/VRChat/Tools`

### 2. dump cookie from browser

windows:

```powershell
./yt-dlp.exe --cookies-from-browser chrome --cookies cookies.txt
```

linux:

```bash
# you should install linux yt-dlp for dump cookie
# for example, archlinux:
#   pacman -S yt-dlp
yt-dlp --cookies-from-browser chrome --cookies cookies.txt
```

## 3. rename yt-dlp.exe name to yt-dlp-o.exe

rename `yt-dlp.exe` to `yt-dlp-o.exe`

## 4. download fix-vrchat-youtube-bot-confirm.exe and rename to yt-dlp.exe

## want use --cookies-from-browser instead of --cookies

It't possible in windows, just create `browser.txt` with you browser name as content, then delete `cookies.txt`

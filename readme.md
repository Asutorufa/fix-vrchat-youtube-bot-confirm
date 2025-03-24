#

This is a yt-dlp wrap which try pass cookies to fix `ERROR: [youtube]: Sign in to confirm you’re not a bot.` and `A web request exception occurred while loading string from URL 'https://youtu.be/xxxxx'. Exception: Redirect limit exceeded` errors in vrchat/VRC.  
It's maybe support in the future versions of vrchat, see: <https://feedback.vrchat.com/open-beta/p/1534-random-error-on-youtube-videos-please-sign-in>.  

## yt-dlp wiki

<https://github.com/yt-dlp/yt-dlp/wiki/FAQ#how-do-i-pass-cookies-to-yt-dlp>

## how to use

### 1. go to vr chat yt-dlp directory

open powershell or linux terminal.  
On windows you can `shift+context-menu` in explorer.  

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

### 3. rename yt-dlp.exe name to yt-dlp-o.exe

rename `yt-dlp.exe` to `yt-dlp-o.exe`.  

```bash
mv yt-dlp.exe yt-dlp-o.exe
```

### 4. download fix-vrchat-youtube-bot-confirm.exe and rename to yt-dlp.exe

windows:

```powershell
curl.exe https://github.com/Asutorufa/fix-vrchat-youtube-bot-confirm/releases/download/v0.0.1/yt-dlp.exe -o yt-dlp.exe
```

linux, you need install `curl` first:

```bash
curl https://github.com/Asutorufa/fix-vrchat-youtube-bot-confirm/releases/download/v0.0.1/yt-dlp.exe -o yt-dlp.exe
```

## want to use --cookies-from-browser instead of --cookies?

It's possible in windows, just create `browser.txt` with you browser name as content, then delete `cookies.txt`

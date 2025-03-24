package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"slices"
	"strings"
)

func main() {
	mydir := filepath.Dir(os.Args[0])

	bin := "yt-dlp-o"
	if runtime.GOOS == "windows" {
		bin += ".exe"
	}

	path := filepath.Join(mydir, bin)
	_, err := os.Stat(path)
	if err != nil {
		path, err = exec.LookPath(strings.Replace(bin, "-o", "", -1))
		if err != nil {
			log.Fatal("yt-dlp not found")
		}
	}

	var args = getCookies(mydir)

	cmd := exec.Command(path, append(args, os.Args[1:]...)...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func getCookies(mydir string) []string {
	if slices.Contains(os.Args, "--cookies") || slices.Contains(os.Args, "--cookies-from-browser") {
		return nil
	}

	cookies := filepath.Join(mydir, "cookies.txt")

	if _, err := os.Stat(cookies); err == nil {
		return []string{
			"--cookies", filepath.Join(mydir, "cookies.txt"),
		}
	}

	data, err := os.ReadFile(filepath.Join(mydir, "browser.txt"))
	if err != nil {
		return nil
	}

	return []string{
		"--cookies-from-browser", string(data),
	}
}

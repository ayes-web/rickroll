package main

import (
	"log"
	"os/exec"
	"runtime"
)

func loo(t string) {
	var r error
	r = exec.Command(t, "https://files.catbox.moe/ok50ut.mp4").Start()
	if r != nil {
		log.Fatal(r)
	}
}
func main() {
	switch runtime.GOOS {
	case "linux":
		loo("xdg-open")
	case "windows":
		loo("start")
	case "darwin":
		loo("open")
	default:
		log.Fatal("unsupported platform")
	}
}

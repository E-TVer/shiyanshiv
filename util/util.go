package util

import (
	"fmt"
	"runtime"
	"os/exec"
)
var commands = map[string]string{
	"darwin":  "open",
	"linux":   "xdg-open",
}
func Info(msg string) {
	fmt.Println("   ")
	fmt.Println("fysys---->"+msg)
}
// Open calls the OS default program for uri
func Open(uri string) error {
	if runtime.GOOS=="windows"{
		cmd := exec.Command("rundll32", "url.dll,FileProtocolHandler", uri)
		return cmd.Start()
	}else {
		run, ok := commands[runtime.GOOS]
		if !ok {
			return fmt.Errorf("don't know how to open things on %s platform", runtime.GOOS)
		}
		cmd := exec.Command(run, uri)
		return cmd.Start()
	}
}
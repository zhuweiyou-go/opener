package openner

import (
	"os/exec"
	"runtime"
)

func Open(args ...string) error {
	var cmd string
	var prefix []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		prefix = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}

	prefix = append(prefix, args...)
	return exec.Command(cmd, prefix...).Start()
}

package utils

import (
	"os"
	"os/user"
	"runtime"
)

func GetHostname() string {
	hostname, err := os.Hostname()
	if err != nil {
		return "unknown-host"
	}
	return hostname
}

func GetUsername() string {
	if currentUser, err := user.Current(); err == nil {
		return currentUser.Username
	}
	if runtime.GOOS == "windows" {
		if username := os.Getenv("USERNAME"); username != "" {
			return username
		}
	} else {
		if username := os.Getenv("USER"); username != "" {
			return username
		}
	}

	return "unknown-user"
}

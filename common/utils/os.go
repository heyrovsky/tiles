package utils

import (
	"fmt"
	"os"
	"os/user"
	"runtime"
	"syscall"

	"golang.org/x/term"
)

func GetHostName() string {
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

func PromptForSecureEntry(message string) (string, error) {
	fmt.Println(message)
	entry, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	if err != nil {
		return "", err
	}
	return string(entry), nil
}

package utils

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
)

func GetSSHAuthWithPassphrase(location, passphrase string) (*ssh.PublicKeys, error) {
	keyPath, err := resolveSSHKeyPath(location)
	if err != nil {
		return nil, err
	}
	privateKey, err := os.ReadFile(keyPath)
	if err != nil {
		return nil, err
	}

	auth, err := createSSHAuth(privateKey, passphrase)
	if err == nil {
		return auth, nil
	}

	if isEncryptedKeyError(err) && passphrase == "" {
		passphrase, err = PromptForSecureEntry(fmt.Sprintf("Enter passphrase for SSH key %s: ", keyPath))
		if err != nil {
			return nil, err
		}
		auth, err = createSSHAuth(privateKey, passphrase)
		if err != nil {
			return nil, err
		}

		return auth, nil
	}
	return nil, err
}

func createSSHAuth(key []byte, passphrase string) (*ssh.PublicKeys, error) {
	auth, err := ssh.NewPublicKeys("git", key, passphrase)
	if err != nil {
		return nil, err
	}
	return auth, nil
}

func resolveSSHKeyPath(location string) (string, error) {
	if location != "" {
		return AbsPath(location)
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("cannot determine user home directory: %w", err)
	}

	sshDir := filepath.Join(homeDir, ".ssh")
	entries, err := os.ReadDir(sshDir)
	if err != nil {
		return "", fmt.Errorf("cannot read ~/.ssh directory: %w", err)
	}

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}
		path := filepath.Join(sshDir, entry.Name())

		// Basic heuristic: check if file starts with typical private key headers
		content, err := os.ReadFile(path)
		if err != nil {
			continue
		}

		if isLikelyPrivateKey(content) {
			return path, nil
		}
	}

	return "", errors.New("no usable SSH key found; set SSH_KEY_LOCATION or place a private key in ~/.ssh/")
}

func isLikelyPrivateKey(content []byte) bool {
	return bytes.HasPrefix(content, []byte("-----BEGIN OPENSSH PRIVATE KEY-----")) ||
		bytes.HasPrefix(content, []byte("-----BEGIN RSA PRIVATE KEY-----")) ||
		bytes.HasPrefix(content, []byte("-----BEGIN EC PRIVATE KEY-----")) ||
		bytes.HasPrefix(content, []byte("-----BEGIN DSA PRIVATE KEY-----"))
}

func isEncryptedKeyError(err error) bool {
	msg := strings.ToLower(err.Error())
	return strings.Contains(msg, "encrypted") ||
		strings.Contains(msg, "passphrase") ||
		strings.Contains(msg, "empty password") ||
		strings.Contains(msg, "decrypt")
}

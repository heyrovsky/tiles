package gitutils

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/go-git/go-git/v5/plumbing/transport/ssh"
	"github.com/heyrovsky/tiles/config"
	"golang.org/x/crypto/ssh/terminal"
)

// GetSSHAuth returns SSH authentication using a passphrase from config or user input if needed.
func GetSSHAuth() (*ssh.PublicKeys, error) {
	return GetSSHAuthWithPassphrase("")
}

// GetSSHAuthWithPassphrase returns SSH authentication using a provided passphrase.
func GetSSHAuthWithPassphrase(passphrase string) (*ssh.PublicKeys, error) {
	keyPath, err := resolveSSHKeyPath()
	if err != nil {
		return nil, err
	}

	privateKey, err := os.ReadFile(keyPath)
	if err != nil {
		return nil, fmt.Errorf("unable to read SSH key from '%s': %w", keyPath, err)
	}

	// Prefer config passphrase
	if config.SSH_KEY_PASS != "" {
		passphrase = config.SSH_KEY_PASS
	}

	auth, err := createSSHAuth(privateKey, passphrase)
	if err == nil {
		return auth, nil
	}

	// Try prompting for passphrase if the error suggests an encrypted key
	if isEncryptedKeyError(err) && passphrase == "" {
		passphrase, perr := promptForPassphrase(keyPath)
		if perr != nil {
			return nil, perr
		}
		auth, err = createSSHAuth(privateKey, passphrase)
		if err != nil {
			return nil, fmt.Errorf("unable to authenticate using provided passphrase: %w", err)
		}
		return auth, nil
	}

	return nil, fmt.Errorf("SSH auth failed: %w", err)
}

// resolveSSHKeyPath returns either the user-defined or the first available default SSH key.
func resolveSSHKeyPath() (string, error) {
	if config.SSH_KEY_LOCATION != "" {
		return config.SSH_KEY_LOCATION, nil
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("cannot determine user home directory: %w", err)
	}

	candidates := []string{
		filepath.Join(homeDir, ".ssh", "id_rsa"),
		filepath.Join(homeDir, ".ssh", "id_ed25519"),
		filepath.Join(homeDir, ".ssh", "id_ecdsa"),
	}

	for _, path := range candidates {
		if _, err := os.Stat(path); err == nil {
			return path, nil
		}
	}

	return "", errors.New("no usable SSH key found; set SSH_KEY_LOCATION or place a private key in ~/.ssh/")
}

// createSSHAuth initializes ssh.PublicKeys, optionally using a passphrase.
func createSSHAuth(key []byte, passphrase string) (*ssh.PublicKeys, error) {
	auth, err := ssh.NewPublicKeys("git", key, passphrase)
	if err != nil {
		return nil, err
	}
	return auth, nil
}

// isEncryptedKeyError determines if an error is related to an encrypted private key.
func isEncryptedKeyError(err error) bool {
	msg := strings.ToLower(err.Error())
	return strings.Contains(msg, "encrypted") ||
		strings.Contains(msg, "passphrase") ||
		strings.Contains(msg, "empty password") ||
		strings.Contains(msg, "decrypt")
}

// promptForPassphrase prompts the user to enter a passphrase securely.
func promptForPassphrase(keyPath string) (string, error) {
	fmt.Printf("Enter passphrase for SSH key %s: ", keyPath)
	password, err := terminal.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	if err != nil {
		return "", fmt.Errorf("failed to read passphrase: %w", err)
	}
	return string(password), nil
}

package utils

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func Command(command string, args ...string) *exec.Cmd {
	cmd := exec.Command(command, args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow: true,
	}
	return cmd
}

func InstallCert(certpath string) error {
	if _, err := os.Stat(certpath); os.IsNotExist(err) {
		return fmt.Errorf("installCert: Certificate does not exist")
	}

	// Register certificate with the Trusted Root store
	cmd := Command("certutil.exe", "-f", "-addstore", "Root", certpath)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("Failed to install Trusted Root certificate: %+v", err)
	}

	// Register certificate with the Trusted Publisher store
	cmd = Command("certutil.exe", "-f", "-addstore", "TrustedPublisher", certpath)
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("Failed to install Trusted Publisher certificate: %+v", err)
	}
	return nil
}

func UninstallCert(authorityName string) error {
	// Remove certificate from TrustedPublisher
	cmd := Command("certutil.exe", "-f", "-delstore", "TrustedPublisher", authorityName)
	err := cmd.Run()
	if err != nil {
		return err
	}

	// Remove certificate from Trusted Root
	cmd = Command("certutil.exe", "-f", "-delstore", "Root", authorityName)
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

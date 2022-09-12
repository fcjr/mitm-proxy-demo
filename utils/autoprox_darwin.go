package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

func EnableProxy(port int) error {
	hwPort, err := getDefaultHWPort()
	if err != nil {
		return err
	}

	// Enable Proxy
	httpEnableCmd := Command("/usr/sbin/networksetup",
		"-setwebproxy",
		hwPort,
		"127.0.0.1",
		strconv.Itoa(port))
	if err := httpEnableCmd.Run(); err != nil {
		return err
	}
	httpsEnableCmd := Command("/usr/sbin/networksetup",
		"-setsecurewebproxy",
		hwPort,
		"127.0.0.1",
		strconv.Itoa(port))
	if err := httpsEnableCmd.Run(); err != nil {
		return err
	}
	return nil
}

func DisableProxy() error {
	hwPort, err := getDefaultHWPort()
	if err != nil {
		return err
	}

	// Disable Proxy
	httpDisableCmd := Command("/usr/sbin/networksetup",
		"-setwebproxystate",
		hwPort,
		"off")
	if err := httpDisableCmd.Run(); err != nil {
		return err
	}
	httpsDisableCmd := Command("/usr/sbin/networksetup",
		"-setsecurewebproxystate",
		hwPort,
		"off")
	if err := httpsDisableCmd.Run(); err != nil {
		return err
	}
	return nil
}

func getDefaultHWPort() (port string, err error) {
	// Get Default Interface
	defaultInterface, err := getDefaultInterface()
	if err != nil {
		return "", err
	}

	getHWPort := Command("networksetup", "-listallhardwareports")
	hwPortBytes, err := getHWPort.Output()
	if err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(bytes.NewReader(hwPortBytes))
	var prevLine string
	for scanner.Scan() {
		line := scanner.Text()
		currTokens := strings.Split(strings.Trim(line, " "), " ")
		for _, token := range currTokens {
			if token == defaultInterface {
				prevLineTokens := strings.Split(strings.Trim(prevLine, " "), ":")
				if len(prevLineTokens) == 2 {
					return strings.Trim(prevLineTokens[1], " "), nil
				}
				return "", fmt.Errorf("failed to parse networksetup output")
			}
		}
		prevLine = line
	}
	return "", fmt.Errorf("failed to parse networksetup output")
}

func getDefaultInterface() (string, error) {
	getInterface := Command("/sbin/route", "get", "default")
	interfaceOut, err := getInterface.Output()
	if err != nil {
		return "", err
	}

	scanner := bufio.NewScanner(bytes.NewReader(interfaceOut))
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "interface") {
			tokens := strings.Split(strings.Trim(line, " "), " ")
			if len(tokens) == 2 {
				return string(tokens[1]), nil
			}
		}
	}
	return "", fmt.Errorf("failed to parse route output")
}

package utils

import "fmt"

func EnableProxy(port int) error {
	// set info
	// set-itemproperty 'HKCU:\Software\Microsoft\Windows\CurrentVersion\Internet Settings' -name ProxyServer -value host:port
	setCmd := Command("powershell", fmt.Sprintf("set-itemproperty 'HKCU:\\Software\\Microsoft\\Windows\\CurrentVersion\\Internet Settings' -name ProxyServer -value 127.0.0.1:%d", port))
	if err := setCmd.Run(); err != nil {
		return err
	}

	// enable proxy
	// set-itemproperty 'HKCU:\Software\Microsoft\Windows\CurrentVersion\Internet Settings' -name ProxyEnable -value 1
	enableCmd := Command("powershell", "set-itemproperty 'HKCU:\\Software\\Microsoft\\Windows\\CurrentVersion\\Internet Settings' -name ProxyEnable -value 1")
	if err := enableCmd.Run(); err != nil {
		return err
	}
	return nil
}

func DisableProxy() error {
	// Disable Proxy
	// set-itemproperty 'HKCU:\Software\Microsoft\Windows\CurrentVersion\Internet Settings' -name ProxyEnable -value 0
	disableCmd := Command("powershell", "set-itemproperty 'HKCU:\\Software\\Microsoft\\Windows\\CurrentVersion\\Internet Settings' -name ProxyEnable -value 0")
	if err := disableCmd.Run(); err != nil {
		return err
	}
	return nil
}

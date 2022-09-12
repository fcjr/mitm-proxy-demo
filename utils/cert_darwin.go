package utils

import (
	"fmt"
	"os/exec"
	"syscall"
	"unsafe"
)

/*
#cgo LDFLAGS: -framework Security -framework CoreFoundation
#include <stdlib.h>
#include <stdio.h>
#include <Security/SecCertificate.h>
#include <Security/SecItem.h>
#include <sys/sysctl.h>
#include <errno.h>

int uninstall_cert(char *authorityName)
{
	CFStringRef str = CFStringCreateWithCString(NULL,
		authorityName, kCFStringEncodingUTF8);
	if (!str) {
		fprintf(stderr, "Error CFStringCreateWithCString");
		return 1;
	}

	const int nkeys = 3;
	CFStringRef keys[nkeys] = {
		kSecClass,
		kSecMatchSubjectWholeString,
		kSecMatchLimit,
	};
	CFTypeRef values[nkeys] = {
		kSecClassCertificate,
		str,
		kSecMatchLimitAll,
	};
	CFDictionaryRef query = CFDictionaryCreate(NULL,
		(const void **)keys, (const void **)values, nkeys,
		NULL, NULL);
	if (!query) {
		fprintf(stderr, "Error CFDictionaryCreate");
		CFRelease(str);
		return 1;
	}

	OSStatus s = SecItemDelete(query);

	CFRelease(query);
	CFRelease(str);
	return s;
}
*/
import "C"

const certMode = syscall.S_IFREG | 0644

func Command(command string, args ...string) *exec.Cmd {
	return exec.Command(command, args...)
}

func InstallCert(certPath string) error {
	var s syscall.Stat_t
	err := syscall.Lstat(certPath, &s)
	if err != nil {
		return err
	}

	if s.Mode != certMode {
		return fmt.Errorf("cert incorrect mode: %o", s.Mode)
	}

	cmd := Command("/usr/bin/security",
		"add-trusted-cert",
		"-d",
		"-k", "/Library/Keychains/System.keychain",
		"-r", "trustRoot",
		certPath)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func UninstallCert(authorityName string) error {
	cAuthorityName := C.CString(authorityName)
	defer C.free(unsafe.Pointer(cAuthorityName))

	ret := C.uninstall_cert(cAuthorityName)
	if ret == C.errSecItemNotFound {
		return fmt.Errorf("sec item not found")
	} else if ret != 0 {
		return fmt.Errorf("uninstall_cert: %d", ret)
	}
	return nil
}

package utils

import (
	"crypto/x509"
	"encoding/pem"
	"os"
)

// WriteCertToFile writes a cert signed with an rsa key to a file on disk
func WriteCertToFile(cert *x509.Certificate, certFilePath string) error {
	// open cert file
	certOut, err := os.OpenFile(certFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer certOut.Close()

	// convert cert to pem format
	certBlock := &pem.Block{
		Type:  "CERTIFICATE",
		Bytes: cert.Raw,
	}

	// write cert to file
	if err := pem.Encode(certOut, certBlock); err != nil {
		os.Remove(certFilePath)
		return err
	}

	// close file
	if err := certOut.Close(); err != nil {
		os.Remove(certFilePath)
		return err
	}

	return nil

}

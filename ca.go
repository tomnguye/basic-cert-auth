package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"fmt"
	"math/big"
	"os"
	"time"
)

func SelfSign() ([]byte, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	publicKey := &privateKey.PublicKey
	if err != nil {
		return nil, fmt.Errorf("error: private key generation failed %w\n", err)
	}
	serialNumberLimit := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, err := rand.Int(rand.Reader, serialNumberLimit)
	if err != nil {
		return nil, fmt.Errorf("error: random number generation failed %w\n", err)
	}
	// generate cert
	template := x509.Certificate{
		SerialNumber:          serialNumber,
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageCRLSign,
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0),
		IsCA:                  true,
		BasicConstraintsValid: true,
	} // key usage, not before not after, constraintsValid, isCA
	certificate, err := x509.CreateCertificate(rand.Reader, &template, &template, publicKey, privateKey)
	if err != nil {
		return nil, fmt.Errorf("error: certificate generation failed %w\n", err)
	}
	fmt.Println(certificate)
	return certificate, nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Print("usage: ca <common name>")
		return
	}
	var selfCertPath = "selfCert.crt"
	_, err := os.ReadFile(selfCertPath)
	if err != nil {
		return // create new cert
	}
	common := os.Args[1]
	fmt.Println(common)
}

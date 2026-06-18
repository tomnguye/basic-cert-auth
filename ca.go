package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
	"os"
)

func SelfSign() {
	// create private key
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: private key generation failed %v\n", err)
		os.Exit(1)
	}
	fmt.Println(privateKey)
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

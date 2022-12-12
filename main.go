package main

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func main() {
	fmt.Println("**************************\nParsePKCS1 RSA Private Key\n**************************")
	fmt.Println("\nThis script decodes a PEM block from a generated RSA key in PKCS1 format, and then parses the key")

	fmt.Println("\nReading RSA key from file....")
	pubPEMData, _ := os.ReadFile("rsa")

	fmt.Println("\nDecoding PEM block....")
	keyPem, _ := pem.Decode(pubPEMData)
	key, err := x509.ParsePKCS1PrivateKey(keyPem.Bytes)

	if err != nil {
		fmt.Println("Could not parse RSA private key")
	}

	fmt.Println("\nParsed PKCS1 key:", key)
}

package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

// GenerateRSAKey generates an RSA key of the specified size (in bits)
func GenerateRSAKey(keySize int) error {
	privateKey, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		return fmt.Errorf("failed to generate RSA key: %v", err)
	}

	err = saveRSAKey(privateKey)
	if err != nil {
		return fmt.Errorf("failed to generate RSA key: %v", err)
	}
	return nil
}

// SaveRSAKey saves the RSA private key to a file
func saveRSAKey(privateKey *rsa.PrivateKey) error {
	keyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	keyPem := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: keyBytes,
	}

	file, err := os.Create("./id_rsa")
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	err = pem.Encode(file, keyPem)
	if err != nil {
		return fmt.Errorf("failed to encode key to PEM: %v", err)
	}

	keyBytes = x509.MarshalPKCS1PublicKey(&privateKey.PublicKey)
	keyPem = &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: keyBytes,
	}

	file, err = os.Create("./id_rsa.pub")
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	err = pem.Encode(file, keyPem)
	if err != nil {
		return fmt.Errorf("failed to encode key to PEM: %v", err)
	}

	return nil
}

// EncryptData encrypts the input data using the RSA public key
func EncryptData(data []byte) ([]byte, error) {
	keyFile, err := os.ReadFile("./id_rsa.pub")
	if err != nil {
		return nil, fmt.Errorf("failed to read key file: %v", err)
	}

	keyPem, _ := pem.Decode(keyFile)
	if keyPem == nil {
		return nil, fmt.Errorf("failed to decode PEM block containing private key")
	}

	publicKey, err := x509.ParsePKCS1PublicKey(keyPem.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse RSA private key: %v", err)
	}

	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, data)
	if err != nil {
		return nil, fmt.Errorf("encryption error: %v", err)
	}

	return ciphertext, nil
}

// DecryptData decrypts the input data using the RSA private key
func DecryptData(ciphertext []byte) ([]byte, error) {
	keyFile, err := os.ReadFile("./id_rsa")
	if err != nil {
		return nil, fmt.Errorf("failed to read key file: %v", err)
	}

	keyPem, _ := pem.Decode(keyFile)
	if keyPem == nil {
		return nil, fmt.Errorf("failed to decode PEM block containing private key")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(keyPem.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse RSA private key: %v", err)
	}

	plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
	if err != nil {
		return nil, fmt.Errorf("decryption error: %v", err)
	}
	return plaintext, nil
}

// GetPublicKey returns the generated RSA public key
func GetPublicKey() ([]byte, error) {
	keyFile, err := os.ReadFile("./id_rsa.pub")
	if err != nil {
		return nil, fmt.Errorf("failed to read key file: %v", err)
	}

	return keyFile, nil
}

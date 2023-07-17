package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"fmt"
	"strconv"
	"time"
)

const AndroidPrefix = "android"

func GenerateAndroidDeviceID() string {
	timestamp := strconv.FormatInt(time.Now().UnixMicro(), 10)
	hash := sha256.Sum256([]byte(timestamp))
	deviceID := hex.EncodeToString(hash[:16])[:16]
	return fmt.Sprintf("%s-%s", AndroidPrefix, deviceID)
}

func getRandomBytes(length int) ([]byte, error) {
	randomBytes := make([]byte, length)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}
	return randomBytes, nil
}

// EncryptPassword encrypt password for authentication
func EncryptPassword(password string, pubKeyID int, pubKeyPem []byte) (string, string, error) {
	passwordAsBytes := []byte(password)

	currentTimestamp := time.Now().Unix()
	currentTimestampAsString := []byte(fmt.Sprintf("%d", currentTimestamp))

	// Generate random secret key and initialization vector
	secretKey, err := getRandomBytes(32)
	if err != nil {
		return "", "", err
	}
	initializationVector, err := getRandomBytes(12)
	if err != nil {
		return "", "", err
	}

	block, _ := pem.Decode(pubKeyPem)
	if block == nil {
		return "", "", fmt.Errorf("failed to parse PEM block containing the public key")
	}
	instagramPublicKeyRaw, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return "", "", err
	}

	instagramPublicKey, ok := instagramPublicKeyRaw.(*rsa.PublicKey)
	if !ok {
		return "", "", errors.New("unknown type of public key")
	}

	// Encrypt secret key using Instagram public key
	encryptedSecretKey, err := rsa.EncryptPKCS1v15(rand.Reader, instagramPublicKey, secretKey)

	if err != nil {
		return "", "", err
	}

	// Encrypt password using secret key and initialization vector
	aesBlock, err := aes.NewCipher(secretKey)
	if err != nil {
		return "", "", err
	}
	aesGCM, err := cipher.NewGCM(aesBlock)
	if err != nil {
		return "", "", err
	}
	encryptedPassword := aesGCM.Seal(nil, initializationVector, passwordAsBytes, currentTimestampAsString)
	authTag := encryptedPassword[len(encryptedPassword)-aesGCM.Overhead():]
	encryptedPassword = encryptedPassword[:len(encryptedPassword)-aesGCM.Overhead()]

	// Construct the password encryption sequence
	passwordEncryptionSequence := make([]byte, 0)

	keyIDMixedBytes := []byte{
		1, byte(0xff & pubKeyID),
	}
	encryptedRSAKeyMixedBytes := []byte{0, 1}

	passwordEncryptionSequence = append(passwordEncryptionSequence, keyIDMixedBytes...)           // Key ID
	passwordEncryptionSequence = append(passwordEncryptionSequence, initializationVector...)      // Initialization Vector
	passwordEncryptionSequence = append(passwordEncryptionSequence, encryptedRSAKeyMixedBytes...) // Encrypted RSA key mixed bytes
	passwordEncryptionSequence = append(passwordEncryptionSequence, encryptedSecretKey...)        // Encrypted secret key
	passwordEncryptionSequence = append(passwordEncryptionSequence, authTag...)                   // Encrypted tag
	passwordEncryptionSequence = append(passwordEncryptionSequence, encryptedPassword...)         // Encrypted password

	passwordAsEncryptionSequenceAsBase64 := base64.StdEncoding.EncodeToString(passwordEncryptionSequence)

	return passwordAsEncryptionSequenceAsBase64, fmt.Sprintf("%d", currentTimestamp), nil
}

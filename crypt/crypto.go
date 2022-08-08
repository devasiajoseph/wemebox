package crypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	b64 "encoding/base64"
	"encoding/hex"
	"errors"
	"io"
	"log"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func RandomString(l int) (string, error) {
	key := make([]byte, l)

	_, err := rand.Read(key)
	if err != nil {
		// handle error here
		log.Panic("Error generating random key")
	}
	return b64.StdEncoding.EncodeToString(key), err
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Fatalf("Hashing password failed")
	}
	return string(bytes)
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func EncryptData(data string, password string) ([]byte, error) {
	dataBytes := []byte(data)
	h := sha256.New()
	h.Write([]byte(password))
	hashKey := h.Sum(nil)

	c, err := aes.NewCipher(hashKey)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	return gcm.Seal(nonce, nonce, dataBytes, nil), nil
}

func DecryptData(encData string, password string) ([]byte, error) {
	cipherText, err := hex.DecodeString(encData)
	if err != nil {
		log.Println("Unable to decode cipher text")
		return nil, err
	}
	h := sha256.New()
	h.Write([]byte(password))
	hashKey := h.Sum(nil)

	c, err := aes.NewCipher(hashKey)
	if err != nil {
		log.Println("Unable to cipher key")
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		log.Println("GCM error")
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(cipherText) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, cipherText := cipherText[:nonceSize], cipherText[nonceSize:]
	return gcm.Open(nil, nonce, cipherText, nil)
}

const otpChars = "1234567890"

//NumericKey generates a numeric key for verification
func NumericKey(length int) (string, error) {
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	otpCharsLength := len(otpChars)
	for i := 0; i < length; i++ {
		buffer[i] = otpChars[int(buffer[i])%otpCharsLength]
	}

	return string(buffer), nil
}

//UID generates unique ID
func UID() string {
	uid := uuid.New()
	return uid.String()
}

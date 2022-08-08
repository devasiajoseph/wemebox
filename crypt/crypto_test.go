package crypt

import (
	"encoding/hex"
	"log"
	"testing"
)

func TestHashPassword(t *testing.T) {
	log.Printf("Testing password hashing")
	phb := HashPassword("password")
	log.Println(phb)
	if !CheckPasswordHash("password", phb) {
		t.Errorf("Password checking failed")
	}

	log.Printf("Testing password hashing done")
}

func TestRandomKey(t *testing.T) {
	log.Printf("Testing random string")
	key, err := RandomString(64)
	if err != nil {
		t.Errorf("Random string failed")
	}
	if len(key) < 64 {
		t.Errorf("Random string length failed")
	}
	log.Printf("Testing random string done")
}

func TestEncryption(t *testing.T) {
	data := "Hello world"
	password := "password"
	enc, err := EncryptData(data, password)
	if err != nil {
		t.Errorf("Encryption failed")
	}
	log.Println("Encrypted data = " + hex.EncodeToString(enc))

	dec, err := DecryptData(hex.EncodeToString(enc), password)
	if err != nil {
		t.Errorf("Decryption failed")
	}
	stringDec := string(dec)
	log.Println("Decrypted data = " + stringDec)
	if data == stringDec {
		log.Println("Successfully decrypted")
	} else {
		t.Errorf("Decryption mismatch")
	}

}

func TestKeyPair(t *testing.T) {
	err := CreateKeys("private.key")
	if err != nil {
		t.Errorf("Key generation failed")
	}
}

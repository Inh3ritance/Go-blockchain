package main

import (
	"crypto/aes"
	"encoding/hex"
)

// Encryption AES
func EncryptAES(key []byte, plaintext string) string {

	c, err := aes.NewCipher(key)
	CheckError(err)

	out := make([]byte, len(plaintext))

	c.Encrypt(out, []byte(plaintext))

	return hex.EncodeToString(out)
}

// Decryption AES
func DecryptAES(key []byte, ct string) string {
	ciphertext, err := hex.DecodeString(ct)

	c, err := aes.NewCipher(key)
	CheckError(err)

	pt := make([]byte, len(ciphertext))
	c.Decrypt(pt, ciphertext)

	s := string(pt[:])
	return s
}

// Checks for errors in Encr/Decr
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

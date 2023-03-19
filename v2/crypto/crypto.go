// Package crypto useful crypto functions
// Reference: https://medium.com/@jcox250/password-hash-salt-using-golang-b041dc94cb72
package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"log"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

// EncryptPwd encrypt password
func EncryptPwd(pwd string) string {
	bytes := []byte(pwd)

	hash, err := bcrypt.GenerateFromPassword(bytes, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		return pwd
	}

	return string(hash)
}

// ComparePwd compare password
func ComparePwd(hashedPwd string, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	bytePlainPwd := []byte(plainPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePlainPwd)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

// EncodeCFB encode CFB Int64
func EncodeCFBInt64(i int64, key string) (string, error) {
	return EncodeCFB(strconv.FormatInt(i, 10), key)
}

// EncodeCFB encode CFB
func EncodeCFB(s string, key string) (string, error) {
	keyBytes := []byte(key)
	plaintext := []byte(s)
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	cipher.NewCFBEncrypter(block, iv).XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
	result := base64.StdEncoding.EncodeToString(ciphertext)
	return result, nil
}

// DecodeCFB decode CFB
func DecodeCFB(s, key string) (string, error) {
	keyBytes := []byte(key)
	ciphertext, _ := base64.StdEncoding.DecodeString(s)
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}
	plaintext := make([]byte, len(ciphertext)-aes.BlockSize)
	iv := ciphertext[:aes.BlockSize]
	cipher.NewCFBDecrypter(block, iv).XORKeyStream(plaintext, ciphertext[aes.BlockSize:])
	return string(plaintext), nil
}

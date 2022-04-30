// Package crypto useful crypto functions
// Reference: https://medium.com/@jcox250/password-hash-salt-using-golang-b041dc94cb72
package crypto

import (
	"log"

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

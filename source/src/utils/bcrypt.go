package utils

import (
    "golang.org/x/crypto/bcrypt"
)

func CryptSomeShit(pwdStr string) (string, error) {
	pwd := []byte(pwdStr)
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	
	return string(hash), nil
}

func CompareSomeShit(hashedPwd string, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	bytePwd := []byte(plainPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, bytePwd)
	
	return err == nil
}
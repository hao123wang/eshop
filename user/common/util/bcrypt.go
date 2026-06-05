package util

import "golang.org/x/crypto/bcrypt"

func Encryption(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func IsCorrectPwd(hashPwd, pwd string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(pwd)); err != nil {
		return false
	}
	return true
}

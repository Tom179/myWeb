package encrypt

import "golang.org/x/crypto/bcrypt"

func EncryptPassword(psw string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(psw), 12)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

package database

import (
	"golang.org/x/crypto/bcrypt"
)

//EncriptarPassword encripta la contraseña para guardarla
func EncriptarPassword(pass string) (string, error) {
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err

}

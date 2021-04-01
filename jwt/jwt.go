package jwt

import (
	"time"

	"github.com/LilJade/virtualBriefcase/models"
	jwt "github.com/dgrijalva/jwt-go"
)

//GeneroJWT genero el encriptado con JWT
func GeneroJwt(t models.Usuario) (string, error) {

	miClave := []byte("Coffee_Encoding")

	//lista de privilegios que se guardan en el payload
	payload := jwt.MapClaims{
		"_id":              t.ID.Hex(),
		"nombre":           t.Nombre,
		"apellidos":        t.Apellidos,
		"direccion":        t.Direccion,
		"fecha_nacimiento": t.FechaNacimiento,
		"email":            t.Email,
		"biografia":        t.Biografia,
		"avatar":           t.Avatar,
		"profesion":        t.Profesion,
		"creado":           t.Creado,
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	//parametro de retorno y error
	tokenStr, err := token.SignedString(miClave)

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil

}

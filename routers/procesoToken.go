package routers

import (
	"errors"
	"strings"

	"github.com/LilJade/virtualBriefcase/database"
	"github.com/LilJade/virtualBriefcase/models"
	"github.com/dgrijalva/jwt-go"
)

var Email string

var IDusuario string

//ProcesoToken proceso token para extraer sus valores
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("Coffee_Encoding")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")

	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}
	tk = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(tk, claims, func(t *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		_, encontrado, _ := database.ChequeoYaExisteUsuario(claims.Email)
		if encontrado == true {
			Email = claims.Email
			IDusuario = claims.ID.Hex()
		}
		return claims, encontrado, IDusuario, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}
	return claims, false, string(""), err
}

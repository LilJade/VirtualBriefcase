package routers

import (
	"encoding/json"
	"net/http"

	"github.com/LilJade/virtualBriefcase/database"
	"github.com/LilJade/virtualBriefcase/jwt"
	"github.com/LilJade/virtualBriefcase/models"
)

//Funcion Login
func ULogin(response http.ResponseWriter, r *http.Request) {
	response.Header().Add("Content-Type", "application/json")

	var usuario models.Usuario

	err := json.NewDecoder(r.Body).Decode(&usuario)

	if err != nil {
		http.Error(response, "Su Corre o contraseña no son invalidos"+err.Error(), 400)
		return
	}
	// mediante len me devuele  un valor y solisito correo
	if len(usuario.Email) == 0 {
		http.Error(response, "Su correo es requido", 400)
		return
	}
	//Verifico si el usuario esta en la base de datos
	usuario, existe := database.Checklogin(usuario.Email, usuario.Password)

	if existe == false {
		http.Error(response, "Correo o contraseña no validos", 400)
		return
	}

	jwtKey, err := jwt.GeneroJwt(usuario)
	if err != nil {
		http.Error(response, "Ocurrio un error al intentar generar el Token correspondiente"+err.Error(), 400)
		return
	}

	resp := models.RespuestaLogin{
		Token: jwtKey,
	}
	response.Header().Set("Content-type", "application/json")
	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(resp)

}

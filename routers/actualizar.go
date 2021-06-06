package routers

import (
	"encoding/json"
	"github.com/LilJade/virtualBriefcase/database"
	"github.com/LilJade/virtualBriefcase/models"
	"net/http"
)



//Metodo para Actualizar en base de datos
func ActualizarDatos(response http.ResponseWriter,request *http.Request)  {
	var Usuario models.Usuario

	err := json.NewDecoder(request.Body).Decode(&Usuario)

	if err !=nil {
		http.Error(response,"datos invalidos"+err.Error(),400)
		return
	}

	var status bool

	status,err=database.ActualizarUsuario(Usuario,IDusuario)


	if err !=nil {
		http.Error(response,"Error al actualizar recargue la pagina e intente de nuevo"+err.Error(),400)
		return
	}

	if status == false {
		http.Error(response,"no se ha logrado actualizar usuario",400)
		return
	}
	response.WriteHeader(http.StatusCreated)

}
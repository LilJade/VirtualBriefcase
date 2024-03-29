package routers

import (
	"net/http"

	"github.com/LilJade/virtualBriefcase/database"
	"github.com/LilJade/virtualBriefcase/models"
)

// AltaRelacionHU es la funcion que se encarga de dar de alta una relacion entre un usuario y una herramienta
func AltaRelacionHU(response http.ResponseWriter, request *http.Request) {
	ID := request.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(response, "es necesario un Id para poder realizar esta accion ", http.StatusBadRequest)
		return
	}

	var h models.RelacionesHu
	h.UsuarioID = IDusuario
	h.HerramientasID = ID

	status, err := database.RelacionHu(h)

	if err != nil {
		http.Error(response, "Error en la accion"+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(response, "No se ha logrado insertar ...", http.StatusBadRequest)
		return
	}

	response.WriteHeader(http.StatusCreated)
}

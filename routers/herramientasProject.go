package routers

import (
	"github.com/LilJade/virtualBriefcase/database"
	"github.com/LilJade/virtualBriefcase/models"
	"net/http"
)

func Utools (response http.ResponseWriter, request *http.Request) {
	ID := request.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(response, "es necesario el ID para realizar esta accion", http.StatusBadRequest)
		return
	}

	var u models.RelacionesHP
	u.ProyectoID=IDusuario
	u.HerramientasID=ID

	status, err := database.ToolsP(u)
	if err != nil {
		http.Error(response, "Error en la accion"+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(response, "No se ha logrado insertar la relacion...", http.StatusBadRequest)
		return
	}

	response.WriteHeader(http.StatusCreated)
}

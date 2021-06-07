package routers

import (
	"encoding/json"
	db "github.com/LilJade/virtualBriefcase/database"
	m "github.com/LilJade/virtualBriefcase/models"
	"net/http"
)

func ActualizarProyecto(w http.ResponseWriter, r *http.Request) {
	var up m.Proyecto

	err := json.NewDecoder(r.Body).Decode(&up)

	if err != nil {
		http.Error(w, "Datos inválidos: "+err.Error(), 400)
		return
	}

	update := m.GraboProyecto{
		UserID:      IDusuario,
		Titulo:      up.Titulo,
		Descripcion: up.Descripcion,
		Portada:     up.Portada,
		Empresa:     up.Empresa,
	}

	var status bool
	status, err = db.ActualizarProyecto(update, IDusuario)

	if err != nil {
		http.Error(w, "error al intentar actualizar el proyecto: "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se pudo completar la actualización...", 400)
		return
	}
}

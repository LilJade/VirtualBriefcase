package routers

import (
	"encoding/json"
	"net/http"
	"time"

	db "github.com/LilJade/virtualBriefcase/database"
	m "github.com/LilJade/virtualBriefcase/models"
)

func InsertoProyecto(w http.ResponseWriter, r *http.Request) {
	var p m.Proyecto
	err := json.NewDecoder(r.Body).Decode(&p)

	proyect := m.GraboProyecto{
		UserID:      IDusuario,
		Titulo:      p.Titulo,
		Github:      p.Github,
		SitioWeb:    p.SitioWeb,
		Descripcion: p.Descripcion,
		Portada:     p.Portada,
		Empresa:     p.Empresa,
		Fecha:       time.Now(),
	}

	_, status, err := db.InsertoProyecto(proyect)

	if err != nil {
		http.Error(w, "Ocurrió un error al intentar registrar el proyecto: "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se pudo insertar el proyecto: ", 400)
	}

	w.WriteHeader(http.StatusCreated)

}

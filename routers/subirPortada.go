package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/LilJade/virtualBriefcase/database"
	"github.com/LilJade/virtualBriefcase/models"
)

func SubirPortada(w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")

	file, handler, err := r.FormFile("portada")
	var extension = strings.Split(handler.Filename, ".")[1]
	var archivo string = "uploads/portada/" + id + "." + extension

	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		http.Error(w, "Error al subir la imagen: "+err.Error(), 400)
		return
	}

	_, err = io.Copy(f, file)

	if err != nil {
		http.Error(w, "Error al copiar la imagen: "+err.Error(), 400)
		return
	}

	var project models.GraboProyecto
	var status bool

	project.Portada = id + "." + extension

	status, err = database.ActualizarProyecto(project, id)

	if err != nil {
		http.Error(w, "No se pudo guardar en la base de datos: "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se pudo guardar en la base de datos", 400)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

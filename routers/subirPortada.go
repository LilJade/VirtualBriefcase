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

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Se necesita el parámetro id", http.StatusBadRequest)
		return
	}
	file, handler, err := r.FormFile("portada")
	var extension = strings.Split(handler.Filename, ".")[1]
	var archivo string = "uploads/portada/" + ID + "." + extension

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

	project.Portada = ID + "." + extension

	status, err = database.ActualizarProyecto(project, ID)

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

package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/LilJade/virtualBriefcase/database"
	"github.com/LilJade/virtualBriefcase/models"
)

func Subiricono(w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")
	file, handler, err := r.FormFile("icono")

	var extension = strings.Split(handler.Filename, ".")[1]
	var archivo string = "uploads/icono/" + id + "." + extension
	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error al subir la imagen !"+err.Error(), http.StatusBadRequest)
		return
	}
	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error al copiar la imagen ! "+err.Error(), http.StatusBadRequest)
		return
	}

	var herramientas models.Herramientas
	var status bool
	herramientas.Imagen = id + "." + extension
	status, err = database.ActualizarHabilidades(herramientas, id)
	if err != nil || status == false {
		http.Error(w, "Error al grabar el avatar en la BD ! "+err.Error(), http.StatusBadRequest)
		return

	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

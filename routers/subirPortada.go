package routers

import (
	"github.com/LilJade/virtualBriefcase/database"
	"github.com/LilJade/virtualBriefcase/models"
	"io"
	"net/http"
	"os"
	"strings"
)

func SubirPortada(w http.ResponseWriter, r *http.Request) {

	file, handler, err := r.FormFile("portada")
	var extension = strings.Split(handler.Filename, ".")[1]
	var archivo string = "uploads/portada/" + IDusuario + "." + extension

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

	project.Portada = IDusuario + "." + extension

	status, err = database.ActualizarProyecto(project, IDusuario)

	if err != nil {
		http.Error(w, "No se pudo guardar en la base de datos: "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se pudo guardar en la base de datos", 400)
		return
	}
}

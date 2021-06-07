package routers

import (
	"github.com/LilJade/virtualBriefcase/database"
	"io"
	"net/http"
	"os"
)

func ObtenerPortada(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Se necesita el ID: ", 400)
		return
	}

	project, err := database.BuscarProyecto(ID)

	if err != nil {
		http.Error(w, "Proyecto no encontrado"+err.Error(), 400)
		return
	}

	OpenFile, err := os.Open("uploads/portada/" + project.Portada)

	if err != nil {
		http.Error(w, "Imagen no encontrada: "+err.Error(), 400)
		return
	}

	_, err = io.Copy(w, OpenFile)

	if err != nil {
		http.Error(w, "Error al copiar la imagen: "+err.Error(), 400)
		return
	}
}

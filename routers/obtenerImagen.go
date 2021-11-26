package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/LilJade/virtualBriefcase/database"
)

func GetImagen(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	IDp := r.URL.Query().Get("proyect")
	if len(ID) < 1 {
		http.Error(w, "Se necesita el ID: ", 400)
		return
	}

	imagen, err := database.BuscarImagen(ID)

	if err != nil {
		http.Error(w, "imagen no encontrado"+err.Error(), 400)
		return
	}

	OpenFile, err := os.Open("uploads/proyectos/" + IDp + "/" + imagen.Imagen)

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

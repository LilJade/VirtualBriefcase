package routers

import (
	"net/http"

	db "github.com/LilJade/virtualBriefcase/database"
)

func BorroImagen(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Se necesita el ID del proyecto", 400)
	}

	err := db.BorroImagen(ID)

	if err != nil {
		http.Error(w, "Ocurrió un error al intentar borrar la imagen: "+err.Error(), 400)
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

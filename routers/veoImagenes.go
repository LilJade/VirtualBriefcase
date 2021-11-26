package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	db "github.com/LilJade/virtualBriefcase/database"
)

func Veoimagenes(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Se necesita el Id", 400)
		return
	}

	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Se necesita la pagina", 400)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("pagina"))

	if err != nil {
		http.Error(w, "El parámetro pagina debe ser > 0", 400)
		return
	}

	pag := int64(page)

	respuesta, correcto := db.VeoImagenes(ID, pag)

	if correcto == false {
		http.Error(w, "Error al leer las imagenes: ", 400)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(respuesta)
}

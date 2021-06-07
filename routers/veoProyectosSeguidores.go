package routers

import (
	"encoding/json"
	"github.com/LilJade/virtualBriefcase/database"
	"net/http"
	"strconv"
)

func VeoProyectosSeguidores(w http.ResponseWriter, r *http.Request) {

	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviar el parametro pagina", 400)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("pagina"))

	if err != nil {
		http.Error(w, "Se necesita la variable pagina como entero mayor a 0", 400)
		return
	}

	respuesta, correcto := database.VeoProyectosSeguidores(IDusuario, page)

	if correcto == false {
		http.Error(w, "Error al mostrar los proyectos", 400)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(respuesta)
}

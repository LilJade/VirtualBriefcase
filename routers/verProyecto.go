package routers

import (
	"encoding/json"
	"net/http"

	"github.com/LilJade/virtualBriefcase/database"
)

//verProyecto para visualizar un proyecto
func VerProyecto(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Se necesita el parámetro id", http.StatusBadRequest)
		return
	}

	proyecto, err := database.BuscarProyecto(ID)

	if err != nil {
		http.Error(w, "Ocurrió un error al consultar el registro..."+err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(proyecto)

}

package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/LilJade/virtualBriefcase/database"
)

//ListaHerramientas es la estructura que contiene la lista de herramientas
func LeoHerramientas(response http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	id := r.URL.Query().Get("id")
	search := ""

	pageTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(response, "deber enviar el parametro pagina mayor a 0 y numero entero", http.StatusBadRequest)
		return
	}
	pag := int64(pageTemp)

	results, status := database.VeoHerramienta(id, pag, search)

	if status == false {
		http.Error(response, "error al leer datos del usuario", http.StatusBadRequest)
		return
	}
	response.Header().Set("Content-type", "application/json")
	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(results)
}

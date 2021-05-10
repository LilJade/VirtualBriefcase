package routers

import (
	"encoding/json"
	"github.com/LilJade/virtualBriefcase/database"
	"net/http"
	"strconv"
)

func Listausuarios(response http.ResponseWriter, r *http.Request)  {
	typeUser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pageTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(response, "deber enviar el parametro  mayor a 0 y numero entero", http.StatusBadRequest)
		return
	}
		pag :=int64(pageTemp)

		results, status :=database.MostrarUsers(IDusuario,pag,search,typeUser)

		if status== false{
			http.Error(response,"erro al leer datos del usuario",http.StatusBadRequest)
			return
		}
		response.Header().Set("Content-type", "application/json")
	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(results)
}
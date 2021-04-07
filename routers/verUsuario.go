package routers

import (
	"encoding/json"
	"github.com/LilJade/virtualBriefcase/database"
	"net/http"
)

//Función para mostrar en Json los datos del usuario
func VerUsuario(w http.ResponseWriter, r *http.Request) {

	//Obtenemos el ID desde la URL
	ID := r.URL.Query().Get("id")

	//Si la variable está vacía devolvemos un error
	if len(ID) < 1 {
		http.Error(w, "Se necesita el parámetro id", http.StatusBadRequest)
		return
	}

	//Si no hay ningun problema, buscamos el ID en la DB
	perfil, err := database.BuscarPerfil(ID)

	//Si al consultar la DB ocurre alun problema, lo mostramos
	if err != nil {
		http.Error(w, "Ocurrió un error al consultar el registro..."+err.Error(), http.StatusBadRequest)
	}

	//Caso contrario enviamos la cabecera
	//y escribimos el Json con los datos del usuario
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(perfil)
}

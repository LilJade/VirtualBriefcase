package routers

import (
	"encoding/json"
	"net/http"

	"github.com/LilJade/virtualBriefcase/database"
	"github.com/LilJade/virtualBriefcase/models"
)

//ActualizarPassword es la funcion para editar en la BD el password del usuario
func ActualizarPassword(w http.ResponseWriter, r *http.Request) {

	var Usuario models.Usuario
	err := json.NewDecoder(r.Body).Decode(&Usuario)

	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}
	var status bool
	status, err = database.ActualizarContrasena(Usuario, IDusuario)

	if err != nil {
		http.Error(w, "Error al actualizar recargue la pagina e intente de nuevo"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "no se ha logrado actualizar usuario", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

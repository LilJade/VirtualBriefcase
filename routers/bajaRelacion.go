package routers

import (
	"net/http"

	"github.com/LilJade/virtualBriefcase/database"
	"github.com/LilJade/virtualBriefcase/models"
)

func BajaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	var t models.Relaciones
	t.UsuarioID = IDusuario
	t.UsuarioRelacionID = ID

	status, err := database.BorroRelacion(t)

	if err != nil {
		http.Error(w, "ocurrio un error al intentar borrar la relacion"+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "no se ha logrado borrar la relacion"+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

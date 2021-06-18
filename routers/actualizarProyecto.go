package routers

import (
	"encoding/json"
	"fmt"
	"net/http"

	db "github.com/LilJade/virtualBriefcase/database"
	m "github.com/LilJade/virtualBriefcase/models"
)

func ActualizarProyecto(w http.ResponseWriter, r *http.Request) {
	var Proyecto m.GraboProyecto

	err := json.NewDecoder(r.Body).Decode(&Proyecto)

	if err != nil {
		http.Error(w, "datos invalidos"+err.Error(), 400)
		return
	}

	var status bool

	status, err = db.ActualizarProyecto(Proyecto, Proyecto.ID.Hex())
	fmt.Println(Proyecto.ID.String())
	fmt.Println(Proyecto.ID.Hex())
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

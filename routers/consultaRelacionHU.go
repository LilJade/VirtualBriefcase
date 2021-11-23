package routers

import (
	"encoding/json"
	"net/http"

	"github.com/LilJade/virtualBriefcase/database"
	"github.com/LilJade/virtualBriefcase/models"
)

func ConsultaRelacionHU(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	var t models.RelacionesHu
	t.UsuarioID = IDusuario
	t.HerramientasID = ID

	var resp models.RespuestaConsultaRelacionHU

	status, err := database.ConsultaRelacionH(t)
	if err != nil || status == false {
		resp.Status = false
	} else {
		resp.Status = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}

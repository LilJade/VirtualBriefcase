package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/LilJade/virtualBriefcase/database"
)

//ObtenerIcono obtiene el icono de habilidad
func ObtenerIcono(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	herramientas, err := database.BuscarHerramienta(ID)
	if err != nil {
		http.Error(w, "Habilidad no encontrado", http.StatusBadRequest)
		return
	}

	OpenFile, err := os.Open("uploads/icono/" + herramientas.Imagen)

	if err != nil {
		http.Error(w, "Imagen no encontrado", http.StatusBadRequest)
		return
	}
	_, err = io.Copy(w, OpenFile)
	if err != nil {
		http.Error(w, "Error al copiar la imagen", http.StatusBadRequest)
	}
}

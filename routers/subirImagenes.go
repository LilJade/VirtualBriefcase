package routers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/LilJade/virtualBriefcase/database"
	"github.com/LilJade/virtualBriefcase/models"
	"github.com/segmentio/ksuid"
)

func SubirImagenes(w http.ResponseWriter, r *http.Request) {

	id := r.FormValue("id")
	idImagen := ksuid.New().String()
	var carpteta string = "uploads/proyectos/" + id

	if _, error := os.Stat(carpteta); os.IsNotExist(error) {
		error = os.Mkdir(carpteta, 0755)
		fmt.Println("si se hizo la carpeta")
		if error != nil {
			panic(error)
		}
	}

	file, handler, erre := r.FormFile("imagen")
	var extension = strings.Split(handler.Filename, ".")[1]
	var archivo string = "uploads/proyectos/" + id + "/" + idImagen + "." + extension

	f, erre := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)

	if erre != nil {
		http.Error(w, "Error al subir la imagen: "+erre.Error(), 400)
		return
	}
	_, err := io.Copy(f, file)

	if err != nil {
		http.Error(w, "Error al copiar la imagen: "+err.Error(), 400)
		return
	}

	var imagen models.GraboImagen
	var status bool

	imagen.ProyectoID = id
	imagen.Imagen = idImagen + "." + extension
	imagen.Fecha = time.Now()

	_, status, err = database.InsertoImagen(imagen)

	if err != nil {
		http.Error(w, "No se pudo guardar en la base de datos: "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se pudo guardar en la base de datos", 400)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

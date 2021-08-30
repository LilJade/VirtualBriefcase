package database_test

import (
	"github.com/LilJade/virtualBriefcase/database"
	"github.com/LilJade/virtualBriefcase/models"
	"testing"
	"time"
)

func TestInsertoProyecto(t *testing.T) {
	proyecto := models.GraboProyecto{
		UserID:      "hola",
		Titulo:      "hjkkk",
		Github:      "khk",
		SitioWeb:    "hgkhh",
		Descripcion: "ssss",
		Portada:     "dsdssd",
		Empresa:     "empresa",
		Fecha:       time.Now(),
	}
	_, _, err := database.InsertoProyecto(proyecto)
	if err != nil {
		t.Error("no se inserto el proyecto")
		t.Fail()
	} else {
		t.Log("si se inserto el proyecto")
	}
}

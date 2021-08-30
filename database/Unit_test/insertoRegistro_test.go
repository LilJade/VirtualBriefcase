package database_test

import (
	"github.com/LilJade/virtualBriefcase/database"
	"github.com/LilJade/virtualBriefcase/models"
	"testing"
)

func TestInsertoRegistro(t *testing.T) {
	usuario := models.Usuario{

		Nombre:    "nom",
		Apellidos: "meo",
	}
	_, _, err := database.InsertoRegistro(usuario)
	if err != nil {
		t.Log("nel no se registro")
		t.Fail()

	} else {
		t.Log("se registro")
	}

}

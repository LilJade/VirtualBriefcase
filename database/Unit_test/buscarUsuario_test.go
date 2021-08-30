package database_test

import (
	"github.com/LilJade/virtualBriefcase/database"
	"testing"
)

func TestBuscarUsuario(t *testing.T) {
	usuario, err := database.BuscarPerfil("60bacf2d8d1f8688854382c3")
	if err != nil {
		t.Error("el usuario no existe")
		t.Fail()

	} else {
		t.Log("el usuario si se encontro")
		t.Log(usuario)
	}

}

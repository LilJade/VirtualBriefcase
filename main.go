package main

import (
	"log"

	db "github.com/LilJade/virtualBriefcase/database"
	"github.com/LilJade/virtualBriefcase/handlers"
)

func main() {
	//Al lanzar la aplicación chequeamos la conexión a la BD
	//Si CheckConnection nos devuelve
	//un cero, la conexión habrá fracasado
	if db.CheckConnection() == 0 {
		log.Fatal("No se pudo conectar con la DB")
		return
	}
	handlers.Manejadores()


}

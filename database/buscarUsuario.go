package database

import (
	"context"
	"fmt"
	"github.com/LilJade/virtualBriefcase/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

//Función para buscar a un usuario mediante su ID
func BuscarPerfil(ID string) (models.Usuario, error) {
	//Establecemos el contexto
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)

	defer cancel()

	//Nos conectamos a la DB y luego a la colección
	db := MongoConn.Database("bd")
	coll := db.Collection("usuario")

	//Declaramos una variable de tipo Usuario
	var perfil models.Usuario
	objId, _ := primitive.ObjectIDFromHex(ID)

	//Establecemos la condición del objeto Id
	condicion := bson.M{
		"_id": objId,
	}

	//Buscamos el Usuario en la DB
	err := coll.FindOne(ctx, condicion).Decode(&perfil)

	//Ocultamos el password
	perfil.Password = ""

	//Si ocurre un error
	if err != nil {
		fmt.Println("Usuario no encontrado!.." + err.Error())
		return perfil, err
	}

	//Si no hay errores, retornamos el usuario encontrado
	return perfil, nil
}

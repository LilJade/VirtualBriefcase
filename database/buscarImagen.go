package database

import (
	"context"
	"fmt"
	"time"

	"github.com/LilJade/virtualBriefcase/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//BusacarHerramienta busca una herramienta en la base de datos
func BuscarImagen(ID string) (models.Imagen, error) {
	//Establecemos el contexto
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)

	defer cancel()

	//Nos conectamos a la DB y luego a la colección
	db := MongoConn.Database("bd")
	coll := db.Collection("ImagenesProyecto")

	//Declaramos una variable de tipo Usuario
	var Imagens models.Imagen
	objId, _ := primitive.ObjectIDFromHex(ID)

	//Establecemos la condición del objeto Id
	condicion := bson.M{
		"_id": objId,
	}

	//Buscamos el Usuario en la DB
	err := coll.FindOne(ctx, condicion).Decode(&Imagens)

	//Si ocurre un error
	if err != nil {
		fmt.Println("imagen no encontrado!.." + err.Error())
		return Imagens, err
	}

	//Si no hay errores, retornamos el usuario encontrado
	return Imagens, nil
}

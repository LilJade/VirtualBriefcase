package database

import (
	"context"

	"time"

	"github.com/LilJade/virtualBriefcase/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ActualizarHabilidades(U models.Herramientas, ID string) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := MongoConn.Database("bd")
	col := database.Collection("herramientas")

	registro := make(map[string]interface{})

	if len(U.Nombre) > 0 {
		registro["nombre"] = U.Nombre
	}
	if len(U.Imagen) > 0 {
		registro["imagen"] = U.Imagen
	}

	updateString := bson.M{ //Convertir el valor bson
		"$set": registro, //
	}

	objID, _ := primitive.ObjectIDFromHex(ID)
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filter, updateString)

	if err != nil {
		return false, err
	}

	return true, nil
}

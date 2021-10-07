package database

import (
	"context"
	"time"

	"github.com/LilJade/virtualBriefcase/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Actualizar el contraseña del usuario

func ActualizarContrasena(U models.Usuario, ID string) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := MongoConn.Database("bd")
	col := database.Collection("usuario")

	contrasenaN := make(map[string]interface{})

	if len(U.Password) > 0 {
		U.Password, _ = EncriptarPassword(U.Password)
		contrasenaN["password"] = U.Password
	}

	updateString := bson.M{ //Convertir el valor bson
		"$set": contrasenaN, //
	}
	objID, _ := primitive.ObjectIDFromHex(ID)

	filter := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filter, updateString)

	if err != nil {
		return false, err
	}

	return true, nil
}

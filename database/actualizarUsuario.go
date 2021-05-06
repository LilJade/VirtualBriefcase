package database

import (
	"context"

	"time"

	"github.com/LilJade/virtualBriefcase/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Actualizar el registro del usuario

func ActualizarUsuario(U models.Usuario, ID string) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	database := MongoConn.Database("bd")
	col := database.Collection("usuario")

	registro := make(map[string]interface{}) // se crea arreglo para  almacenar datos temporalmente

	if len(U.Nombre) > 0 {
		registro["nombre"] = U.Nombre
	}

	if len(U.Apellidos) > 0 {
		registro["apellidos"] = U.Apellidos
	}

	if len(U.Direccion) > 0 {
		registro["direccion"] = U.Direccion
	}

	registro["fechaNacimiento"] = U.FechaNacimiento

	if len(U.Email) > 0 {
		registro["email"] = U.Email
	}
	if len(U.Avatar) > 0 {
		registro["avatar"] = U.Avatar
	}
	if len(U.Biografia) > 0 {
		registro["biografia"] = U.Biografia
	}
	if len(U.Profesion) > 0 {
		registro["profesion"] = U.Profesion
	}

	registro["creado"] = U.Creado

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

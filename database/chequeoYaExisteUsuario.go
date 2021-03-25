package database

import (
	"context"
	"time"

	"github.com/LilJade/virtualBriefcase/models"
	"go.mongodb.org/mongo-driver/bson"
)

//ChequeoYaExisteUsuario recibe un email y chequea si exite en la base de datos
func ChequeoYaExisteUsuario(email string) (models.Usuario, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	db := MongoConn.Database("bd")
	col := db.Collection("usuario")

	condicion := bson.M{"email": email}

	var resultado models.Usuario

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()
	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID

}

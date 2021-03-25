package database

import (
	"context"
	"time"

	"github.com/LilJade/virtualBriefcase/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//InsertoRegistro es la parada fina con la BD para insertar los datos del usuario
func InsertoRegistro(u models.Usuario) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConn.Database("bd")
	col := db.Collection("usuario")

	u.Password, _ = EncriptarPassword(u.Password)

	result, err := col.InsertOne(ctx, u)

	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil

}

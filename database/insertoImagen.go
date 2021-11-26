package database

import (
	"context"
	"time"

	"github.com/LilJade/virtualBriefcase/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//InsertoImagen inserta una imagen en la base de datos
func InsertoImagen(t models.GraboImagen) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	database := MongoConn.Database("bd")
	col := database.Collection("ImagenesProyecto")

	registro := bson.M{
		"proyectoid": t.ProyectoID,
		"titulo":     t.Titulo,
		"imagen":     t.Imagen,
		"fecha":      t.Fecha,
	}

	result, err := col.InsertOne(ctx, registro)

	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil

}

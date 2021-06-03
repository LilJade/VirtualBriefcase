package database

import (
	"context"
	m "github.com/LilJade/virtualBriefcase/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func InsertoProyecto(p m.GraboProyecto) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConn.Database("bd")
	col := db.Collection("proyecto")

	registro := bson.M{
		"userid":      p.UserID,
		"titulo":      p.Titulo,
		"descripcion": p.Descripcion,
		"portada":     p.Portada,
		"empresa":     p.Empresa,
		"fecha":       p.Fecha,
	}

	result, err := col.InsertOne(ctx, registro)

	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)

	return objID.String(), true, nil

}

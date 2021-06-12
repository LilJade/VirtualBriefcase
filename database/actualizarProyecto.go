package database

import (
	"context"
	"time"

	m "github.com/LilJade/virtualBriefcase/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ActualizarProyecto(P m.GraboProyecto, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConn.Database("bd")
	col := db.Collection("proyectos")

	datos := make(map[string]interface{})

	if len(P.Titulo) > 0 {
		datos["titulo"] = P.Titulo
	}
	if len(P.Github) > 0 {
		datos["github"] = P.Titulo
	}
	if len(P.SitioWeb) > 0 {
		datos["sitioWeb"] = P.Titulo
	}
	if len(P.Descripcion) > 0 {
		datos["descripcion"] = P.Descripcion
	}

	if len(P.Portada) > 0 {
		datos["portada"] = P.Portada
	}

	if len(P.Empresa) > 0 {
		datos["empresa"] = P.Empresa
	}

	updateString := bson.M{"$set": datos}

	objID, _ := primitive.ObjectIDFromHex(ID)

	filtro := bson.M{"_id": bson.M{"$eq": objID}}

	_, err := col.UpdateOne(ctx, filtro, updateString)

	if err != nil {
		return false, err
	}

	return true, err
}

package database

import (
	"context"
	"fmt"
	"time"

	m "github.com/LilJade/virtualBriefcase/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BuscarProyecto(ID string) (m.GraboProyecto, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConn.Database("bd")
	col := db.Collection("proyectos")

	var project m.GraboProyecto
	objID, _ := primitive.ObjectIDFromHex(ID)

	condicion := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx, condicion).Decode(&project)

	if err != nil {
		fmt.Println("Proyecto no encontrado: " + err.Error())
		return project, err
	}

	return project, nil
}

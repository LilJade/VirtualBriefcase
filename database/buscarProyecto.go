package database

import (
	"context"
	"fmt"
	m "github.com/LilJade/virtualBriefcase/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func BuscarProyecto(ID string) (m.GraboProyecto, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConn.Database("db")
	col := db.Collection("proyecto")

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

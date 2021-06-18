package database

import (
	"context"
	"time"

	m "github.com/LilJade/virtualBriefcase/models"
	"go.mongodb.org/mongo-driver/bson"
)

func VeoProyectosSeguidores(ID string, page int) ([]m.DevuelvoProyectoSeguidores, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConn.Database("bd")
	col := db.Collection("relaciones")

	skip := (page - 1) * 20

	conditions := make([]bson.M, 0)
	conditions = append(conditions, bson.M{"$match": bson.M{"usuarioid": ID}})
	conditions = append(conditions, bson.M{
		"$lookup": bson.M{
			"from":         "proyectos",
			"localField":   "usuariorelacionid",
			"foreignField": "userid",
			"as":           "proyectos",
		}})
	conditions = append(conditions, bson.M{"$unwind": "$proyectos"})
	conditions = append(conditions, bson.M{"$sort": bson.M{"proyectos.fecha": -1}})
	conditions = append(conditions, bson.M{"$skip": skip})
	conditions = append(conditions, bson.M{"$limit": 20})

	cursor, err := col.Aggregate(ctx, conditions)

	var result []m.DevuelvoProyectoSeguidores

	err = cursor.All(ctx, &result)

	if err != nil {
		return result, false
	}

	return result, true
}

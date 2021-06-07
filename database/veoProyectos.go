package database

import (
	"context"
	m "github.com/LilJade/virtualBriefcase/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func VeoProyectos(ID string, page int64) ([]*m.DevuelvoProyecto, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConn.Database("bd")
	col := db.Collection("proyecto")

	var result []*m.DevuelvoProyecto

	condition := bson.M{
		"userid": ID,
	}

	opciones := options.Find()
	opciones.SetLimit(20)
	opciones.SetSort(bson.D{{Key: "fecha", Value: -1}})
	opciones.SetSkip((page - 1) * 20)

	cursor, err := col.Find(ctx, condition, opciones)

	if err != nil {
		log.Fatal(err.Error())
		return result, false
	}

	for cursor.Next(context.TODO()) {
		var reg m.DevuelvoProyecto
		err := cursor.Decode(&reg)

		if err != nil {
			return result, false
		}

		result = append(result, &reg)
	}

	return result, true
}

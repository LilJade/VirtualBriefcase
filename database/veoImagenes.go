package database

import (
	"context"
	"log"
	"time"

	m "github.com/LilJade/virtualBriefcase/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func VeoImagenes(ID string, page int64) ([]*m.DevuelvoImagenes, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConn.Database("bd")
	col := db.Collection("ImagenesProyecto")

	var result []*m.DevuelvoImagenes

	condition := bson.M{
		"proyectoid": ID,
	}

	opciones := options.Find()
	opciones.SetLimit(9)
	opciones.SetSort(bson.D{{Key: "fecha", Value: -1}})
	opciones.SetSkip((page - 1) * 9)

	cursor, err := col.Find(ctx, condition, opciones)

	if err != nil {
		log.Fatal(err.Error())
		return result, false
	}

	for cursor.Next(context.TODO()) {
		var reg m.DevuelvoImagenes
		err := cursor.Decode(&reg)

		if err != nil {
			return result, false
		}

		result = append(result, &reg)

	}
	return result, true

}

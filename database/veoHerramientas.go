package database

import (
	"context"
	"fmt"
	"time"

	"github.com/LilJade/virtualBriefcase/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MostrarHerramientas: Muestra todas las herramientas de la base de datos
func VeoHerramienta(ID string, page int64, search string) ([]*models.Herramientas, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	bd := MongoConn.Database("bd")
	col := bd.Collection("herramientas")

	var results []*models.Herramientas

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"nombre": bson.M{"$regex": `(?i)` + search},
	}

	cur, err := col.Find(ctx, query, findOptions)
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	var encontrado, incluir bool

	for cur.Next(ctx) {
		var s models.Herramientas
		err := cur.Decode(&s)
		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}

		var r models.RelacionesHu
		r.UsuarioID = ID
		r.HerramientasID = s.ID.Hex()

		incluir = false
		encontrado, err = ConsultaRelacionH(r)

		if encontrado == true {
			incluir = true
		}

		if incluir == true {

			results = append(results, &s)
		}
	}

	err = cur.Err()
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}
	cur.Close(ctx)
	return results, true

}

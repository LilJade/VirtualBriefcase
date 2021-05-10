package database

import (
	"context"
	"fmt"
	"github.com/LilJade/virtualBriefcase/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)


//funcion mostrar usuarios recibe 3 parametros
func  MostrarUsers(ID string, page int64, search string,tipo string) ([]*models.Usuario,bool) {
	ctx,cancel := context.WithTimeout(context.Background(),15*time.Second)
	defer cancel()

	bd :=MongoConn.Database("bd")
	col := bd.Collection("usuario")

	var results []*models.Usuario//Guardamos los datos en un array para poder mostrarlos

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	query :=bson.M{
		"nombre": bson.M{"$regex": `(?i)` + search},
	}

	cur, err := col.Find(ctx, query, findOptions)
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	var  encontrado,incluir bool
	for cur.Next(ctx){
		var s models.Usuario
		err := cur.Decode(&s)
		if err != nil {
			fmt.Println(err.Error())
			return results, false

	}
		incluir=true

		if tipo =="follow" && encontrado==true{
			incluir= true
		}



//datos que no se mostraran
		if incluir == true && tipo=="follow"{


			s.Password=""



			results =append(results,&s)


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


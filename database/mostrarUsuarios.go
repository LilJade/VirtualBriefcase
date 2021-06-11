package database

import (
	"context"
	"fmt"
	"time"

	"github.com/LilJade/virtualBriefcase/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//funcion mostrar usuarios recibe 3 parametros
func MostrarUsers(ID string, page int64, search string, tipo string) ([]*models.Usuario, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	bd := MongoConn.Database("bd")
	col := bd.Collection("usuario")

	var results []*models.Usuario //Guardamos los datos en un array para poder mostrarlos

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
		var s models.Usuario
		err := cur.Decode(&s)
		if err != nil {
			fmt.Println(err.Error())
			return results, false

		}
		var r models.Relaciones
		r.UsuarioID = ID
		r.UsuarioRelacionID = s.ID.Hex()

		incluir = false

		encontrado, err = ConsultaRelacion(r)

		if tipo == "new" && encontrado == false {
			incluir = true
		}

		if tipo == "follow" && encontrado == true {
			incluir = true
		}

		if r.UsuarioRelacionID == ID {
			incluir = false
		}
		//datos que no se mostraran
		if incluir == true {

			s.Password = ""
			s.Biografia = ""
			s.Direccion = ""
			s.Email = ""
			s.Instagram = ""
			s.Facebook = ""
			s.Twitter = ""

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

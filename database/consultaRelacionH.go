package database

import (
	"context"
	"fmt"
	"time"

	"github.com/LilJade/virtualBriefcase/models"
	"go.mongodb.org/mongo-driver/bson"
)

//consultaRelacionH consulta la relacion de una herramienta con un usuario
func ConsultaRelacionH(t models.RelacionesHu) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConn.Database("bd")
	col := db.Collection("relacionHU")

	condicion := bson.M{
		"usuarioid":      t.UsuarioID,
		"herramientasid": t.HerramientasID,
	}

	var resultado models.RelacionesHu
	fmt.Println(resultado)
	err := col.FindOne(ctx, condicion).Decode(&resultado)
	if err != nil {
		fmt.Println(err.Error())
		return false, err
	}
	return true, nil
}

package database

import (
	"context"
	"github.com/LilJade/virtualBriefcase/models"
	"time"
)

func Relacion(u models.Relaciones) (bool, error) {
	ctx,cancel := context.WithTimeout(context.Background(),15*time.Second)
	defer cancel()

	db :=MongoConn.Database("bd")
	col :=db.Collection("relaciones")

	_, err := col.InsertOne(ctx, u)
	if err != nil {
		return false, err
	}
	return true, nil
}
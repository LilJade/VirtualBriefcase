package database

import (
	"context"
	"github.com/LilJade/virtualBriefcase/models"
	"time"
)

func ToolsP(u models.RelacionesHP) (bool,error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db :=MongoConn.Database("bd")
	col :=db.Collection("relacionHP")
	_, err := col.InsertOne(ctx, u)
	if err != nil {
		return false, err
	}
	return true, nil
}
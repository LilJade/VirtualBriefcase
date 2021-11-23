package database

import (
	"context"
	"github.com/LilJade/virtualBriefcase/models"
	"time"
)


func RelacionHu(u models.RelacionesHu)(bool, error)  {
	ctx, cancel := context.WithTimeout(context.Background(),20*time.Second)
	defer cancel()

	db:=MongoConn.Database("bd")
	col :=db.Collection("relacionHU")

	_, err :=col.InsertOne(ctx,u)
	if err !=nil{
		return false,err
	}
	return true, nil

}
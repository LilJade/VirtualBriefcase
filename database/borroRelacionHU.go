package database

import (
	"context"
	"time"

	"github.com/LilJade/virtualBriefcase/models"
)

//BorroRelacion borra la relacion en la BD
func BorroRelacionHU(t models.RelacionesHu) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConn.Database("bd")
	col := db.Collection("relacionHU")

	_, err := col.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}

package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DevuelvoHerramientas struct {
	ID     primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID string             `bson:"userid" json:"userid,omitempty"`
	Nombre string             `bson:"nombre" json:"nombre,omitempty"`
	Imagen string             `bson:"imagen" json:"imagen,omitempty"`
}

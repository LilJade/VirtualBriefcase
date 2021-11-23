package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Herramientas struct {
	ID     primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Nombre string             `bson:"nombre" json:"nombre,omitempty"`
	Imagen string             `bson:"imagen" json:"imagen,omitempty"`
}

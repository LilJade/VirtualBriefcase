package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DevuelvoImagenes struct {
	ID         primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	ProyectoID string             `bson:"proyectoid" json:"proyectoid,omitempty"`
	Imagen     string             `bson:"imagen" json:"imagen,omitempty"`
}

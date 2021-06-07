package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type DevuelvoProyecto struct {
	ID          primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserID      string             `bson:"userid" json:"userid,omitempty"`
	Titulo      string             `bson:"titulo" json:"titulo,omitempty"`
	Descripcion string             `bson:"descripcion" json:"descripcion,omitempty"`
	Portada     string             `bson:"portada" json:"portada,omitempty"`
	Empresa     string             `bson:"empresa" json:"empresa,omitempty"`
	Fecha       time.Time          `bson:"fecha" json:"fecha,omitempty"`
}

/*
	`bson:"" json:",omitempty"`
*/
